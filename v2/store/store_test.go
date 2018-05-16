package store

import (
	"testing"
)

func counter(state int, action string) int {
	switch action {
	case "INC":
		return state + 1
	case "DEC":
		return state - 1
	default:
		return state
	}
}

type withPayload struct {
	typ     string
	payload int
}

func payload(state int, action withPayload) int {
	switch action.typ {
	case "INC":
		return state + action.payload
	case "DEC":
		return state - action.payload
	default:
		return state
	}
}

func TestStoreNew(t *testing.T) {
	store := /*store.*/ New(counter)
	if store == nil {
		t.Error("func New fail, expected it return a pointer to store instance")
	}
}

func TestStoreDispatch(t *testing.T) {
	store := /*store.*/ New(counter, payload)
	store.Dispatch("INC")
	if store.GetState(counter) != 1 {
		t.Error("counter can not work")
	}

	store.Dispatch(withPayload{
		typ:     "INC",
		payload: 10,
	})
	if store.GetState(payload) != 10 {
		t.Error("payload should increase by payload")
	}
}

func TestStoreGetState(t *testing.T) {
	store := /*store.*/ New(counter)
	store.Dispatch("DEC")
	if store.GetState(counter) != -1 {
		t.Error("GetState should return reducer's state")
	}
}

func TestStoreSubscribe(t *testing.T) {
	store := /*store.*/ New(counter)

	store.Subscribe(func() {})

	if len(store.subscribedFuncs) != 1 {
		t.Error("Subscribe do not work")
	}
}
