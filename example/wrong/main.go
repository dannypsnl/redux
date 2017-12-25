package main

import (
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
		store.Dispatch(redux.SendAction("INC"))
		//store.Subscribe(func() {})
	})
	store.Dispatch(redux.SendAction("INC"))
}
