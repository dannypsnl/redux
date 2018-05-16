package store

import (
	"reflect"
)

// Store is a type manage your data
//
// Usage:
//   counter := func(state int, payload int) int {
//       return state + payload
//   }
//   store := store.New(counter)
//   store.Dispatch(30)
//   fmt.Printf("%d\n", store.GetState(counter)) // expected: 30
type Store struct {
	reducers        []reflect.Value
	state           map[uintptr]reflect.Value
	subscribedFuncs []func()
}

// New create a Store by reducers
func New(reducers ...interface{}) *Store {
	newStore := &Store{
		reducers: make([]reflect.Value, 0),
		state:    make(map[uintptr]reflect.Value),
	}
	for _, reducer := range reducers {
		r := reflect.ValueOf(reducer)
		// If fail any checking, it will panic, so don't try to recover or handling the error
		checkReducer(r)
		newStore.reducers = append(newStore.reducers, r)

		newStore.state[r.Pointer()] = r.Call(
			[]reflect.Value{
				// We just use their zero value for initialize
				reflect.Zero(r.Type().In(0)), // In index 0 is state
				reflect.Zero(r.Type().In(1)), // In index 1 is action
			},
		)[0] // 0 at here is because checkReducer promise that we will only receive one return
	}
	return newStore
}

// Dispatch send action to all reducer in Store to update state
//
// In v2, action can be anything you want, it's more powerful rather than v1
func (s *Store) Dispatch(action interface{}) {
	for _, r := range s.reducers {
		if reflect.ValueOf(action).Kind() == r.Type().In(1).Kind() {
			res := r.Call(
				[]reflect.Value{
					s.state[r.Pointer()],
					reflect.ValueOf(action)})
			s.state[r.Pointer()] = res[0]
		}
	}

	for _, s := range s.subscribedFuncs {
		s()
	}
}

// Subscribe let user emit a function will be triggered by Dispatch
func (s *Store) Subscribe(function func()) {
	s.subscribedFuncs = append(s.subscribedFuncs, function)
}

// GetState return the reducer name matches state
func (s *Store) GetState(reducer interface{}) interface{} {
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
