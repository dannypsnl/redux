package redux

import (
	"sync"
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

func BenchmarkAlotEasySubscribe(b *testing.B) {
	n := 0
	store := NewStore(counter, jump)
	for i := 0; i < 100000000; i++ {
		n++
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(SendAction("JUMP"))
	}
}

func Benchmark10000workConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(1 * time.Nanosecond)
			}()
			wg.Wait()
		}
	}
}

func Benchmark10000work(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// If quest cost Nanosecond, sequence is faster than concurrency way.
		// I think the problem is paging
		// It about 3~4 time
		for i := 0; i < 10000; i++ {
			time.Sleep(1 * time.Nanosecond)
		}
	}
}
