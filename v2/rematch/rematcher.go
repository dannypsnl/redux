package rematch

import (
	"reflect"
)

// Reducer is core of rematch
//
// It should be embedded by using side like
//
// Code:
//  type Counter struct {
//      rematch.Reducer
//  }
type Reducer struct {
	// State is not mutable field, it represents initial state
	State interface{}
}

func (r Reducer) Action(method interface{}) *Action {
	return &Action{
		addr: reflect.ValueOf(method).Pointer(),
	}
}

type Action struct {
	addr uintptr
}

func (a Action) Addr() uintptr {
	return a.addr
}

func (a *Action) With(payload interface{}) *Action {
	return a
}
