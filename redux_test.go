package redux

import (
	"testing"
)

func counter(state interface{}, action Action) interface{} {
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

func jump(state interface{}, action Action) interface{} {
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
	store := NewStore(counter, jump)
	store.Subscribe(func() {
		if store.GetState(thisState) != expectedState {
			t.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState(thisState))
		}
	})
	store.Dispatch(SendAction("JUMP"))
	thisState = "counter"
	expectedState = 1
	store.Dispatch(SendAction("INC"))
}

func TestSubscribtorCallSubscribeShouldPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`should panic when subscribetor trying to call store::Subscribe`)
		}
	}()
	store := NewStore(counter)
	store.Subscribe(func() {
		store.Subscribe(func() {})
	})
	store.Dispatch(SendAction("INC"))
}

func TestSubscribtorCallSubscribeByDispatchCShouldPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`should panic when subscribtor trying to call store::Subscribe`)
		}
	}()
	store := NewStore(counter)
	store.Subscribe(func() {
		store.Subscribe(func() {})
	})
	store.DispatchC(SendAction("INC"))
}

func TestDispatchC(t *testing.T) {
	store := NewStore(counter)
	store2 := NewStore(counter)
	store.Dispatch(SendAction("INC"))
	store2.DispatchC(SendAction("INC"))
	if store.GetState("counter") != store2.GetState("counter") {
		t.Errorf(
			"DispatchC & Dispatch are different, Dispatch: %d, DispatchC: %d",
			store.GetState("counter"),
			store2.GetState("counter"),
		)
	}
}
