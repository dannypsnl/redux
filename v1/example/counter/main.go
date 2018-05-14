package main

import (
	"fmt"
	"github.com/dannypsnl/redux/v1/action"
	"github.com/dannypsnl/redux/v1/store"
)

func counter(state interface{}, act action.Action) interface{} {
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

func main() {
	store := store.New(counter)
	store.Subscribe(func() {
		fmt.Println(store.GetState("counter"))
	})
	store.Dispatch(action.New("INC"))
}
