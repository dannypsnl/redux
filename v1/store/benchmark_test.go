package store

import (
	"github.com/dannypsnl/redux/v1/action"
	"testing"
	"time"
)

func Benchmark_2_Reducers_1_Subscribe(b *testing.B) {
	expectedState := 0
	store := New(counter)
	store.Subscribe(func() {
		expectedState++
		if store.GetState("counter") != expectedState {
			b.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState("jump"))
		}
	})
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("INC"))
	}
}

func sleepNnsSubscribe(n int64, b *testing.B) {
	store := New(counter)
	for i := 0; i < 10000; i++ {
		store.Subscribe(func() {
			time.Sleep(time.Duration(n) * time.Nanosecond)
		})
	}
	for i := 0; i < b.N; i++ {
		// Subscribtor can only invoke by Dispatch routine
		store.Dispatch(action.New("INC"))
	}
}

func BenchmarkSleep1nsSubscribe(b *testing.B)     { sleepNnsSubscribe(1, b) }
func BenchmarkSleep175nsSubscribe(b *testing.B)   { sleepNnsSubscribe(175, b) }
func BenchmarkSleep375nsSubscribe(b *testing.B)   { sleepNnsSubscribe(375, b) }
func BenchmarkSleep575nsSubscribe(b *testing.B)   { sleepNnsSubscribe(575, b) }
func BenchmarkSleep775nsSubscribe(b *testing.B)   { sleepNnsSubscribe(775, b) }
func BenchmarkSleep1000nsSubscribe(b *testing.B)  { sleepNnsSubscribe(1000, b) }
func BenchmarkSleep1500nsSubscribe(b *testing.B)  { sleepNnsSubscribe(1500, b) }
func BenchmarkSleep10000nsSubscribe(b *testing.B) { sleepNnsSubscribe(10000, b) }

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
