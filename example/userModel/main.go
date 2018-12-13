package main

import (
	"fmt"

	"github.com/dannypsnl/redux/v2/example/userModel/user"
	"github.com/dannypsnl/redux/v2/store"
)

func main() {
	store := store.New(user.Reducer)
	dan := store.StateOf(user.Reducer).(user.Model)
	fmt.Printf("user name: %s, age: %d\n", dan.Name, dan.Age)

	updateName := user.Reducer.Action(user.Reducer.UpdateName)
	store.Dispatch(updateName.With("Danny"))
	store.Dispatch(user.Reducer.Action(user.Reducer.UpdateAge).With(21))

	danny := store.StateOf(user.Reducer).(user.Model)
	fmt.Printf("After update, user name: %s, age: %d\n", danny.Name, danny.Age)
	// expect: user name: Danny, age: 21
}
