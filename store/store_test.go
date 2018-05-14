package store

import (
	"testing"

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

func counter(state int, action string) int {
	switch action {
	case "INC":
		return state + 1
	case "DEC":
		return state - 1
	default:
		return state
	}
}

func TestStoreNew(t *testing.T) {
	store := /*store.*/ New(counter)
	if store == nil {
		t.Error("func New fail, expected it return a pointer to store instance")
	}
}

func TestStoreDispatch(t *testing.T) {
	store := /*store.*/ New(counter)
	store.Dispatch("INC")
	if store.state["counter"].Interface() != 1 {
		t.Error("counter can not work")
	}
}

func TestStoreGetState(t *testing.T) {
	store := /*store.*/ New(counter)
	store.Dispatch("DEC")
	if store.GetState("counter") != -1 {
		t.Error("GetState should return reducer's state")
	}
}
