package redux

import (
	"testing"
	"time"
)

func BenchmarkEasySubscribe(b *testing.B) {
	store := NewStore(counter, jump)
	store.Subscribe(func() {
		expectedState := "TOP"
		if store.GetState("jump") != expectedState {
			b.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState("jump"))
		}
	})
	store.Subscribe(func() {
		expectedState := "TOP"
		if store.GetState("jump") != expectedState {
			b.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState("jump"))
		}
	})
	for i := 0; i < b.N; i++ {
		store.Dispatch(SendAction("JUMP"))
	}
}

func BenchmarkHeavySubscribe(b *testing.B) {
	store := NewStore(counter, jump)
	store.Subscribe(func() {
		time.Sleep(1 * time.Millisecond)
	})
	store.Subscribe(func() {
		time.Sleep(1 * time.Millisecond)
	})
	for i := 0; i < b.N; i++ {
		store.Dispatch(SendAction("JUMP"))
	}
}
