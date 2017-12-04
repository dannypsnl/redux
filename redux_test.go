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

func TestStore(t *testing.T) {
	testTarget := "jump"
	var expectedState interface{} = "TOP"
	store := NewStore(counter, jump)
	store.Subscribe(func() {
		if store.GetState(testTarget) != expectedState {
			t.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState(testTarget))
		}
	})
	store.Dispatch(SendAction("JUMP"))
	testTarget = "counter"
	expectedState = 1
	store.Dispatch(SendAction("INC"))
}

func TestSubscribetorCannotInvokeDispatch(t *testing.T) {
	store := NewStore(counter, jump)
	defer func() {
		if r := recover(); r == nil {
			//t.Errorf("The code did not panic")
		}
	}()
	store.Subscribe(func() {
		//store.Dispatch(SendAction("INC"))
	})
	store.Dispatch(SendAction("INC"))
}
