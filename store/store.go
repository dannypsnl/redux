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

func New(reducers ...interface{}) *Store {
	rs := make([]reflect.Value, 0)
	stateInit := make(map[string]reflect.Value)
	for _, reducer := range reducers {
		r := reflect.ValueOf(reducer)
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

		stateType := r.Type().In(0)
		rs = append(rs, r)

		reducerName := runtime.FuncForPC(r.Pointer()).Name()
		res := r.Call(
			[]reflect.Value{
				reflect.Zero(stateType),
				reflect.Zero(r.Type().In(1))})
		stateInit[reducerName[strings.LastIndexByte(reducerName, '.')+1:]] = res[0]
	}
	s := &Store{
		reducers: rs,
		state:    stateInit,
	}
	return s
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
