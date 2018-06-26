package main

import (
	"fmt"

	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

type user struct {
	name string
	age  int
}

type UserModel struct {
	rematch.Reducer
	State user
}

func (cm *UserModel) UpdateName(s user, newName string) user {
	return user{
		name: newName,
		age:  s.age,
	}
}

func (cm *UserModel) UpdateAge(s user, newAge int) user {
	return user{
		name: s.name,
		age:  newAge,
	}
}

func main() {
	c := &UserModel{
		State: user{
			name: "dan",
			age:  20,
		},
	}
	store := store.New(c)
	store.Dispatch(c.Action(c.UpdateName).With("Danny"))
	store.Dispatch(c.Action(c.UpdateAge).With(21))

	danny := store.StateOf(c).(user)
	fmt.Printf("user name: %s, age: %d\n", danny.name, danny.age)
	// expect: user name: Danny, age: 21
}
