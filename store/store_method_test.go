package store

import (
	"github.com/dannypsnl/redux/action"
	"testing"
)

func TestStoreState(t *testing.T) {
	thisState := "jump"
	var expectedState interface{} = "TOP"

	store := /*store.*/ New(counter, jump)
	store.Subscribe(func() {
		if store.GetState(thisState) != expectedState {
			t.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState(thisState))
		}
	})
	store.Dispatch(action.New("JUMP"))

	thisState = "counter"
	expectedState = 1

	store.Dispatch(action.New("INC"))
}

func TestSubscribtorCallSubscribeShouldPanic(t *testing.T) {
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
