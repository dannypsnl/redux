package store

import (
	"testing"

	"reflect"
)

type Store struct {
}

func New(reducers ...interface{}) *Store {
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
	}
	return &Store{}
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
