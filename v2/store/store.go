package store

import (
	"reflect"
	"sync"
)

// Store is a type manage your data
//
// Usage:
//   counter := func(state int, payload int) int {
//       return state + payload
//   }
//   store := store.New(counter)
//   store.Dispatch(30)
//   fmt.Printf("%d\n", store.StateOf(counter)) // expected: 30
type Store struct {
	reducers            []reflect.Value
	state               map[uintptr]reflect.Value
	subscribedFuncs     []func()
	dispatch            sync.Mutex
	subscribe           sync.Mutex
	subscribedFuncPanic bool
	onDispatching       bool
}

// New create a Store by reducers
func New(reducers ...interface{}) *Store {
	newStore := &Store{
		reducers: make([]reflect.Value, 0),
		state:    make(map[uintptr]reflect.Value),

		subscribedFuncPanic: false,
	}
	for _, reducer := range reducers {
		r := reflect.ValueOf(reducer)

		if r.Kind() == reflect.Ptr {
			v := reflect.Indirect(r)
			if v.FieldByName("State").Kind() == reflect.Invalid {
				panic("Reducer structure must contains field[State]")
			}

			if _, ok := newStore.state[r.Pointer()]; ok {
				panic("You can't put duplicated reducer into the same store!")
			}

			ms := make(map[uintptr]reflect.Value)
			for i := 1; i < v.NumMethod(); i++ {
				m := v.Method(i)
				if m.Type().NumIn() == 2 &&
					m.Type().NumOut() == 1 &&
					m.Type().In(0) == m.Type().Out(0) {
					ms[m.Pointer()] = m
				}
			}

			insideReducer := func(s interface{}, a interface {
				Addr() uintptr
				Payload() interface{}
			}) interface{} {
				return ms[a.Addr()].Call(
					[]reflect.Value{
						reflect.ValueOf(s),
						reflect.ValueOf(a.Payload()),
					},
				)[0]
			}

			r := reflect.ValueOf(insideReducer)
			newStore.reducers = append(newStore.reducers, r)

			newStore.state[r.Pointer()] = reflect.ValueOf(v.FieldByName("State").Interface())
		} else {
			// If fail any checking, it will panic, so don't try to recover or handling the error
			checkReducer(r)

			if _, ok := newStore.state[r.Pointer()]; ok {
				panic("You can put duplicated reducer into the same store!")
			}

			newStore.reducers = append(newStore.reducers, r)

			newStore.state[r.Pointer()] = r.Call(
				[]reflect.Value{
					// We just use their zero value for initialize
					reflect.Zero(r.Type().In(0)), // In index 0 is state
					reflect.Zero(r.Type().In(1)), // In index 1 is action
				},
			)[0] // 0 at here is because checkReducer promise that we will only receive one return
		}
	}
	return newStore
}

// Dispatch send action to all reducer in Store to update state
//
// In v2, action can be anything you want, it's more powerful rather than v1
func (s *Store) Dispatch(action interface{}) {
	s.dispatch.Lock()
	defer s.dispatch.Unlock()

	rAction := reflect.ValueOf(action)
	for _, r := range s.reducers {
		if rAction.Kind() == r.Type().In(1).Kind() {
			res := r.Call(
				[]reflect.Value{
					s.state[r.Pointer()],
					rAction})
			s.state[r.Pointer()] = res[0]
		}
	}

	s.onDispatching = true

	var wg sync.WaitGroup
	wg.Add(len(s.subscribedFuncs))
	for _, subscribedFunc := range s.subscribedFuncs {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					s.subscribedFuncPanic = true
				}
				wg.Done()
			}()
			subscribedFunc()
		}()
	}
	wg.Wait()

	if s.subscribedFuncPanic {
		panic("You can't call Subscribe in subscribed function")
	}
	s.onDispatching = false
}

// Subscribe let user emit a function will be triggered by Dispatch
func (s *Store) Subscribe(function func()) {
	s.subscribe.Lock()
	defer s.subscribe.Unlock()
	if s.onDispatching {
		panic("You can't call Subscribe in subscribed function")
	}
	s.subscribedFuncs = append(s.subscribedFuncs, function)
}

// StateOf return the reducer name matches state
func (s *Store) StateOf(reducer interface{}) interface{} {
	place := reflect.ValueOf(reducer).Pointer()
	return s.state[place].Interface()
}

// checkReducer reject all unexpected reducer format
func checkReducer(r reflect.Value) {
	if r.Kind() == reflect.Invalid {
		panic("It's an invalid value")
	}

	// reducer :: (state, action) -> state
	if r.Type().NumIn() != 2 {
		panic("reducer should have state & action two parameter, not thing more")
	}
	if r.Type().NumOut() != 1 {
		panic("reducer should return state only")
	}
	if r.Type().In(0) != r.Type().Out(0) {
		panic("reducer should own state with the same type at anytime, if you want have variant value, please using interface")
	}
}
