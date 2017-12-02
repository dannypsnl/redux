package redux

import (
	"testing"
)

func counter(state interface{}, act Action) interface{} {
	// The initial state
	if state == nil {
		return 0
	}
	switch act.Type {
	case "INC":
		return state.(int) + 1
	case "DEC":
		return state.(int) - 1
	default:
		return state
	}
}

func TestStore(t *testing.T) {
	store := NewStore()
	store.NewReducer(counter)
	store.Dispatch(SendAction("INC"))
	if store.GetState["counter"] != 1 {
		t.Errorf("State should be 1 now, but it's: %v", store.GetState["counter"])
	}
}
