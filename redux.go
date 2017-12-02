package redux

import (
	_ "fmt"
	"reflect"
	"runtime"
	"strings"
)

type Action struct {
	Type string
}

func SendAction(typ string) *Action {
	return &Action{
		Type: typ,
	}
}

type Reducer func(interface{}, Action) interface{}

type store struct {
	reducers []Reducer
	// GetState contain key map state
	// and we use key to spread reducer's state
	// for example: "counter" mapping to counter reducer.
	// when use GetState["counter"], we will got the current state of "counter" key's mapping target.
	GetState map[string]interface{}
	// defaultState only work when we don't have several reducer
	defaultState interface{}
}

func NewStore(defaultReducer Reducer, reducers ...Reducer) *store {
	s := &store{
		GetState: make(map[string]interface{}),
	}
	if len(reducers) == 0 {
		return &store{
			defaultState: defaultReducer(nil, Action{}),
			reducers:     []Reducer{defaultReducer},
		}
	}
	s.NewReducer(defaultReducer)
	for _, reducer := range reducers {
		s.NewReducer(reducer)
	}
	return s
}

func (s *store) NewReducer(reducer Reducer) {
	// this code will get package.function_name, so we drop package part
	func_name := runtime.FuncForPC(reflect.ValueOf(reducer).Pointer()).Name()
	func_name = func_name[strings.LastIndexByte(func_name, '.')+1:]
	s.GetState[func_name] = reducer(nil, Action{})
	s.reducers = append(s.reducers, reducer)
}

func (s *store) Dispatch(act *Action) {
	if s.defaultState != nil {
		s.defaultState = s.reducers[0](s.defaultState, *act)
		return
	}
	for _, reducer := range s.reducers {
		func_name := runtime.FuncForPC(reflect.ValueOf(reducer).Pointer()).Name()
		func_name = func_name[strings.LastIndexByte(func_name, '.')+1:]
		s.GetState[func_name] = reducer(s.GetState[func_name], *act)
	}
}
func (s *store) GetStateD() interface{} {
	return s.defaultState
}
