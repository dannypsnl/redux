package store

import (
	"testing"

	"reflect"
)

type Store struct {
}

func New(reducers ...interface{}) *Store {
	for _, r := range reducers {
		reflect.ValueOf(r)
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
