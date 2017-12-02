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
	store := NewStore(counter)
	store.Dispatch(SendAction("INC"))
	if store.GetStateD() != 1 {
		t.Errorf("State should be 1 now, but it's: %v", store.GetStateD())
	}
	store = NewStore(counter, jump)
	store.Dispatch(SendAction("JUMP"))
	if store.GetState["jump"] != "TOP" {
		t.Error(`Error`)
	}
	if store.GetState["counter"] != 0 {
		t.Error(`initial state of counter is wrong`)
	}
}
