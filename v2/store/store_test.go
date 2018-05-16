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
	if store.StateOf(counter) != 1 {
		t.Error("counter can not work")
	}

	store.Dispatch(withPayload{
		typ:     "INC",
		payload: 10,
	})
	if store.StateOf(payload) != 10 {
		t.Error("payload should increase by payload")
	}
}

func TestStoreStateOf(t *testing.T) {
	store := /*store.*/ New(counter)
	store.Dispatch("DEC")
	if store.StateOf(counter) != -1 {
		t.Error("StateOf should return reducer's state")
	}
}

func TestStoreSubscribe(t *testing.T) {
	store := /*store.*/ New(counter)

	store.Subscribe(func() {})

	if len(store.subscribedFuncs) != 1 {
		t.Error("Subscribe do not work")
	}
}

func TestStoreWorkWithLambda(t *testing.T) {
	lambda := func(state int, action int) int {
		return state + action
	}
	store := /*store.*/ New(lambda)
	store.Dispatch(10)
	state := store.StateOf(lambda)
	if state != 10 {
		t.Error("store can't work with lambda")
	}
}

func TestPanic(t *testing.T) {
	error1 := func(state int, act string) string {
		return ""
	}
	error2 := func(state string, act string) (string, error) {
		return "", nil
	}
	error3 := func() {}
	testPanic(t, error1, "should panic when reducer return a state is type different than input state type")
	testPanic(t, error2, "should panic when reducer return several type")
	testPanic(t, error3, "should panic when reducer do not contain state & action two parameters")
	testPanic(t, nil, "should panic for the invalid value")
}

func testPanic(t *testing.T, reducer interface{}, msg string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(msg)
		}
	}()
	New(reducer)
}