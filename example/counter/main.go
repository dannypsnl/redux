package main

import (
	"fmt"
	"github.com/dannypsnl/redux"
)

func counter(state interface{}, act redux.Action) interface{} {
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
	store := redux.NewStore(counter)
	store.Subscribe(func() {
		fmt.Println(store.GetState("counter"))
	})
	store.Dispatch(redux.SendAction("INC"))
}
