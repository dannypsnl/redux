package store

import (
	"github.com/dannypsnl/redux/action"
	"sync"
	"testing"
)

func TestDuplicatedReducerShouldCausePanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`duplicated reducer should break the process`)
		}
	}()
	store := /*store.*/ New(counter, counter)
	store.Dispatch(action.New("INC"))
}

func TestDispatchInConcurrencyIsSafe(t *testing.T) {
	store := /*store.*/ New(counter)
	n := 1000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			store.Dispatch(action.New("INC"))
		}()
	}
	wg.Wait()
	if store.GetState("counter") != 1000 {
		t.Errorf("Expected: %d, Actual: %d", 1000, store.GetState("counter"))
	}
}

func TestStoreStateCanBeUpdateByDispatch(t *testing.T) {
	thisState := "jump"
	expectedState := "TOP"

	store := /*store.*/ New(counter, jump)
	store.Dispatch(action.New("JUMP"))

	if store.GetState(thisState) != expectedState {
		t.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState(thisState))
	}
}

func TestCallSubscribeInSubscribtorShouldPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`should panic when subscribetor trying to call store::Subscribe`)
		}
	}()

	store := /*store.*/ New(counter)
	store.Subscribe(func() {
		store.Subscribe(func() {})
	})

	store.Dispatch(action.New("INC"))
}
