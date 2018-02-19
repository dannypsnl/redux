package store

import (
	"github.com/dannypsnl/redux/action"
	"testing"
	"time"
)

func Benchmark_2_Reducers_1_Subscribe(b *testing.B) {
	store := New(counter, jump)
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

func BenchmarkSleep1nsSubscribe(b *testing.B) {
	store := New(counter, jump)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(1 * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkSleep175nsSubscribe(b *testing.B) {
	store := New(counter, jump)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(175 * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkSleep1msSubscribe(b *testing.B) {
	store := New(counter, jump)
	for i := 0; i < 10; i++ {
		store.Subscribe(func() {
			time.Sleep(1 * time.Millisecond)
		})
	}
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("JUMP"))
	}
}

func BenchmarkStoreUpdateState(b *testing.B) {
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.updateState(action.New("INC"))
	}
}

func (s *Store) updateStateReduceCode(act *action.Action) {
	// dispatch action to every reducer, then reducer update it's mapping state.
	for _, r := range s.reducers {
		funcName := getReducerName(r) // getReducerName in util.go
		s.state[funcName] = r(s.state[funcName], *act)
	}
}

// This one is a little bit slower(usually)
// And more important issue is harder to read
// So keep this solution, make sure everyone won't consider it again
func BenchmarkStoreUpdateStateReduceCode(b *testing.B) {
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.updateStateReduceCode(action.New("INC"))
	}
}
