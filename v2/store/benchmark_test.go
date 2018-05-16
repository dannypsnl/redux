package store

import (
	"testing"

	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/store"
)

func BenchmarkV1StoreDispatch(b *testing.B) {
	counter := func(state interface{}, action action.Action) interface{} {
		if state == nil {
			return 0
		}
		switch action.Type {
		case "INC":
			return state.(int) + 1
		case "DEC":
			return state.(int) - 1
		default:
			return state
		}
	}
	store := store.New(counter)
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("INC"))
	}
}

func BenchmarkStoreDispatch(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.Dispatch(1)
	}
}

func BenchmarkV1StoreGetState(b *testing.B) {
	counter := func(state interface{}, action action.Action) interface{} {
		if state == nil {
			return 0
		}
		switch action.Type {
		case "INC":
			return state.(int) + 1
		case "DEC":
			return state.(int) - 1
		default:
			return state
		}
	}
	store := store.New(counter)
	for i := 0; i < b.N; i++ {
		store.GetState("func1")
	}
}

func BenchmarkStoreStateOf(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.StateOf(counter)
	}
}

func BenchmarkV1StoreSubcribe(b *testing.B) {
	counter := func(state interface{}, action action.Action) interface{} {
		if state == nil {
			return 0
		}
		switch action.Type {
		case "INC":
			return state.(int) + 1
		case "DEC":
			return state.(int) - 1
		default:
			return state
		}
	}
	store := store.New(counter)
	for i := 0; i < b.N; i++ {
		store.Subscribe(func() {})
	}
}

func BenchmarkStoreSubcribe(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.Subscribe(func() {})
	}
}
