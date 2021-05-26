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
	reducers            map[uintptr]reflect.Value
	state               map[uintptr]reflect.Value
	subscribedFuncs     []func()
	dispatch            sync.Mutex
	subscribe           sync.Mutex
	subscribedFuncPanic bool
	onDispatching       bool
}

func getReducerAndInitState(r reflect.Value) (reflect.Value, reflect.Value) {
	if r.Kind() == reflect.Ptr {
		v := reflect.Indirect(r) // dereference from ptr
		return r.MethodByName("InsideReducer").
				Call([]reflect.Value{r})[0],
			v.FieldByName("State")
	}
	return r,
		r.Call(
			[]reflect.Value{
				// We just use their zero value for initialize
				reflect.Zero(r.Type().In(0)), // In index 0 is state
				reflect.Zero(r.Type().In(1)), // In index 1 is action
			},
		)[0] // 0 at here is because checkReducer promise that we will only receive one return
}

// New create a Store by reducers
func New(reducers ...interface{}) *Store {
	newStore := &Store{
		reducers: make(map[uintptr]reflect.Value),
		state:    make(map[uintptr]reflect.Value),
	}
	for _, reducer := range reducers {
		r := reflect.ValueOf(reducer)
		// If fail any checking, it will panic, so don't try to recover or handling the error
		checkReducer(r)

		if _, ok := newStore.state[r.Pointer()]; ok {
			panic("You can't put duplicated reducer into the same store!")
		}

		actualReducer, initState := getReducerAndInitState(r)

		newStore.reducers[r.Pointer()] = actualReducer
		newStore.state[r.Pointer()] = initState
	}
	return newStore
}

// Dispatch send action to reducers in Store to update state
//
// In v2, action can be anything you want
//
// For example:
//   func counter(state, payload int) int {
//   	return state + payload
//   }
//
// Action type should be `int`, if you use others type, just use others type
func (s *Store) Dispatch(action interface{}) {
	s.dispatch.Lock()
	defer s.dispatch.Unlock()

	rAction := reflect.ValueOf(action)
	for at, r := range s.reducers {
		// action type matches second argument
		if rAction.Kind() == r.Type().In(1).Kind() {
			// next state
			s.state[at] = r.Call([]reflect.Value{
				s.state[at], // now state
				rAction,
			})[0]
		}
	}

	s.onDispatching = true

	var wg sync.WaitGroup
	wg.Add(len(s.subscribedFuncs))
	for _, subscribedFunc := range s.subscribedFuncs {
		go func(f func()) {
			defer func() {
				if r := recover(); r != nil {
					s.subscribedFuncPanic = true
				}
				wg.Done()
			}()
			f()
		}(subscribedFunc)
	}
	wg.Wait()

	if s.subscribedFuncPanic {
		panic("You can't call Subscribe in subscribed function")
	}
	s.onDispatching = false
}

// Subscribe emit a function will be triggered after executing Dispatch
func (s *Store) Subscribe(function func()) {
	s.subscribe.Lock()
	defer s.subscribe.Unlock()
	if s.onDispatching {
		panic("You can't call Subscribe in subscribed function")
	}
	s.subscribedFuncs = append(s.subscribedFuncs, function)
}

// StateOf return the reducer matches state
func (s *Store) StateOf(reducer interface{}) interface{} {
	ofReducer := reflect.ValueOf(reducer).Pointer()
	return s.state[ofReducer].Interface()
}

// checkReducer reject all unexpected reducer format
func checkReducer(r reflect.Value) {
	if r.Kind() == reflect.Invalid {
		panic("It's an invalid value")
	}

	if r.Kind() == reflect.Ptr {
		v := reflect.Indirect(r) // dereference from ptr
		if v.FieldByName("State").Kind() == reflect.Invalid {
			panic("Reducer structure must contains field[State]")
		}
	} else {
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
}
