package main

import (
	"fmt"

	rd "github.com/dannypsnl/redux"
)

func counter(state interface{}, action rd.Action) interface{} {
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
	store := rd.NewStore(counter)
	fmt.Printf("State is %v\n", store.GetState())
	store.Dispatch(rd.SendAction("INC"))
	fmt.Printf("State is %v\n", store.GetState())
	store.Dispatch(rd.SendAction("DEC"))
	fmt.Printf("State is %v\n", store.GetState())
}
