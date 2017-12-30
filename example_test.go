package redux_test

import (
	"fmt"

	"github.com/dannypsnl/redux"
	"github.com/dannypsnl/redux/action"
)

func counter(state interface{}, action action.Action) interface{} {
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
	store := redux.NewStore(counter)
	store.Subscribe(func() {
		fmt.Println("Current State:", store.GetState("counter"))
	})
	store.Dispatch(action.New("INC"))
	// Output:
	// Current State: 1
}
