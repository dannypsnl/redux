package redux_test

import (
	"fmt"

	"github.com/dannypsnl/redux"
)

var users = map[string]int{
	// trust me, I won't use 1234 to be my password
	"danny": 1234,
}

func login(state interface{}, action redux.Action) interface{} {
	if state == nil {
		return "Guest"
	}
	switch action.Type {
	case "login":
		if users[action.Args["user"].(string)] == action.Args["password"] {
			return action.Args["user"].(string) + " Login"
		}
		return state
	default:
		return state
	}
}

func ExampleActionArgs() {
	store := redux.NewStore(login)
	act := &redux.Action{
		Type: "login",
		Args: map[string]interface{}{
			"user":     "danny",
			"password": 1234,
		},
	}
	store.Subscribe(func() {
		if store.GetState("login") != "Guest" {
			fmt.Println(store.GetState("login"))
		}
	})
	store.Dispatch(act)
	// Output:
	// danny Login
}
