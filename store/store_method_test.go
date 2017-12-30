package store

import (
	"github.com/dannypsnl/redux/action"
	"testing"
)

func counter(state interface{}, action action.Action) interface{} {
	// The initial state
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

func jump(state interface{}, action action.Action) interface{} {
	if state == nil {
		return "TOP"
	}
	switch action.Type {
	case "JUMP":
		return "TOP"
	case "FALL":
		return "DOWN"
	default:
		return state
	}
}

func TestStoreState(t *testing.T) {
	thisState := "jump"
	var expectedState interface{} = "TOP"
	store := New(counter, jump)
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
	store := New(counter)
	store.Subscribe(func() {
		store.Subscribe(func() {})
	})
	store.Dispatch(action.New("INC"))
}