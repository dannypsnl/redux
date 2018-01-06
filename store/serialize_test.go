package store

import (
	"encoding/json"
	"github.com/dannypsnl/redux/action"
	"testing"
)

func TestSerialize(t *testing.T) {
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

// Make sure Marshal can work with Struct Type
func TestSerializeStruct(t *testing.T) {
	store := /*store.*/ New(fileUpdator)
	expected := []byte(`{"fileUpdator":{"ext":"elz","mod":"+x"}}`)

	if store.Marshal() != string(expected) {
		t.Errorf("expected: %s, actual: %s", expected, store.Marshal())
	}
}

func TestMarshalPanicIfDataIsInvalid(t *testing.T) {
	// At here, we have function in state
	store := /*store.*/ New(foo)
	defer func() {
		if r := recover(); r == nil {
			t.Error(`Marshal didn't panic when data can't Marshal`)
		}
	}()
	store.Marshal()
}
