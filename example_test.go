package redux_test

import (
	"fmt"
	"github.com/dannypsnl/redux"
)

func counter(state interface{}, action redux.Action) interface{} {
	if state == nil {
		return 0
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

func Example() {
	// This Example show how to
	store := redux.NewStore(counter)
	store.Subscribe(func() {
		fmt.Println(store.GetState["counter"])
	})
	store.Dispatch(redux.SendAction("INC"))

	// Output:
	// 1
}
