package store

import (
	"encoding/json"
	"github.com/dannypsnl/redux/v1/action"
	"testing"
)

func TestMarshalAndUnmarshal(t *testing.T) {
	type StoreByCounterLogin struct {
		Counter int    `json:"counter"`
		Login   string `json:"login"`
	}

	store := /*store.*/ New(counter, login)
	store.Subscribe(func() {
		// Use Subscribe to test is the most easy way to reduce lots of code
		expected := []byte(store.Marshal())
		expectedStore := &StoreByCounterLogin{}
		json.Unmarshal(expected, expectedStore)
		if store.GetState("counter") != expectedStore.Counter || store.GetState("login") != expectedStore.Login {
			t.Errorf("serialized result is not expected, expected:\n`%s`, actual:\n`%s`", expected, store.Marshal())
		}
	})

	store.Dispatch(action.New("INC"))
	store.Dispatch(action.New("INC"))
	store.Dispatch(
		action.New("login").
			Arg("user", "danny").
			Arg("password", 1234))
}

func foo(state interface{}, act action.Action) interface{} {
	if state == nil {
		return func() {}
	}
	return state
}

func TestMarshalWillPanicIfDataIsInvalid(t *testing.T) {
	// At here, we have function in state
	store := /*store.*/ &Store{
		state:        make(map[string]interface{}),
		doMiddleware: func(a *action.Action) *action.Action { return a },
	}
	store.state["not_important"] = func() {} // func can't be Marshal
	defer func() {
		if r := recover(); r == nil {
			t.Error(`Marshal didn't panic when data can't Marshal`)
		}
	}()
	store.Marshal()
}
