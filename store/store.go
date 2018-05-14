package store

import (
	"reflect"
	"runtime"
	"strings"
)

// Store is a type manage your data
type Store struct {
	reducers []reflect.Value
	state    map[string]reflect.Value
}

// New create a Store by reducers
//
// Usage:
//   store := store.New(reducer...)
func New(reducers ...interface{}) *Store {
	rs := make([]reflect.Value, 0)
	initState := make(map[string]reflect.Value)
	for _, reducer := range reducers {
		r := reflect.ValueOf(reducer)
		// If fail any checking, it will panic, so don't try to recover or handling the error
		checkReducer(r)
		rs = append(rs, r)

		rName := runtime.FuncForPC(r.Pointer()).Name()
		rName = rName[strings.LastIndexByte(rName, '.')+1:]
		res := r.Call(
			[]reflect.Value{
				reflect.Zero(r.Type().In(0)),
				reflect.Zero(r.Type().In(1))})
		initState[rName] = res[0]
	}
	return &Store{
		reducers: rs,
		state:    initState,
	}
}

// Dispatch send action to all reducer in Store to update state
//
// In v2, action can be anything you want, it's more powerful rather than v1
//
// Usage:
//   store.Dispatch("INC")
func (s *Store) Dispatch(action interface{}) {
	for _, r := range s.reducers {
		rName := runtime.FuncForPC(r.Pointer()).Name()
		rName = rName[strings.LastIndexByte(rName, '.')+1:]
		res := r.Call(
			[]reflect.Value{
				s.state[rName],
				reflect.ValueOf(action)})

		s.state[rName] = res[0]
	}
}

// GetState return the reducer name matches state
func (s *Store) GetState(rName string) interface{} {
	return s.state[rName].Interface()
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
