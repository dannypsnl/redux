package main

import (
	"fmt"

	"github.com/dannypsnl/redux"
)

func counter(state interface{}, action redux.Action) interface{} {
	// This is initial state
	if state == nil {
		state = 0
	}
	switch action.Type {
	case "INC":
		return state.(int) + 1
	case "DEC":
		return state.(int) - 1
	default:
		return state
	}
}

func main() {
	store := redux.NewStore()
	store.NewReducer(counter)
	fmt.Printf("State is %v\n", store.GetState["counter"])
	store.Dispatch(redux.SendAction("INC"))
	fmt.Printf("State is %v\n", store.GetState["counter"])
	store.Dispatch(redux.SendAction("DEC"))
	fmt.Printf("State is %v\n", store.GetState["counter"])
	store.Dispatch(redux.SendAction("DEC"))
	fmt.Printf("State is %v\n", store.GetState["counter"])
}
