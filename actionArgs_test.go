package redux_test

import (
	"fmt"

	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/store"
)

var users = map[string]int{
	// trust me, I won't use 1234 to be my password
	"danny": 1234,
}

func login(state interface{}, action action.Action) interface{} {
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

func Example_actionArgs() {
	store := store.New(login)
	store.Subscribe(func() {
		if store.GetState("login") != "Guest" {
			fmt.Println(store.GetState("login"))
		}
	})
	store.Dispatch(
		action.New("login").
			Arg("user", "danny").
			Arg("password", 1234))
	// Output:
	// danny Login
}
