package store

import (
	"reflect"
	"runtime"
	"strings"
)

type Store struct {
	reducers []reflect.Value
	state    map[string]reflect.Value
}

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

func (s *Store) GetState(rName string) interface{} {
	return s.state[rName].Interface()
}
