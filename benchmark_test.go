package redux

import (
	"github.com/dannypsnl/redux/action"
	"testing"
	"time"
)

func BenchmarkANormalCaseSubscribe(b *testing.B) {
	store := NewStore(counter, jump)
	store.Subscribe(func() {
		expectedState := "TOP"
		if store.GetState("jump") != expectedState {
			b.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState("jump"))
		}
	})
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkSleep1msSubscribe(b *testing.B) {
	store := NewStore(counter, jump)
	for i := 0; i < 10; i++ {
		store.Subscribe(func() {
			time.Sleep(1 * time.Millisecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkSleep1msSubscribeByDispatchC(b *testing.B) {
	store := NewStore(counter, jump)
	for i := 0; i < 10; i++ {
		store.Subscribe(func() {
			time.Sleep(1 * time.Millisecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.DispatchC(action.New("JUMP"))
	}
}

func BenchmarkALotSleep175nsSubscribe(b *testing.B) {
	store := NewStore(counter, jump)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(175 * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkALotSleep175nsSubscribeByDispatchC(b *testing.B) {
	store := NewStore(counter, jump)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(175 * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.DispatchC(action.New("JUMP"))
	}
}

func BenchmarkALotSleep1nsSubscribe(b *testing.B) {
	store := NewStore(counter, jump)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(1 * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkALotSleep1nsSubscribeByDispatchC(b *testing.B) {
	store := NewStore(counter, jump)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(1 * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.DispatchC(action.New("JUMP"))
	}
}
