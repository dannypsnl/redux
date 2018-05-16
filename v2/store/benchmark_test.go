package store

import (
	"testing"
)

func BenchmarkStoreDispatch(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.Dispatch(1)
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

func BenchmarkStoreSubcribe(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.Subscribe(func() {})
	}
}
