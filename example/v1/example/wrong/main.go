package main

import (
	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/store"
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
		store.Dispatch(action.New("INC"))
		// store.Dispatch(redux.SendAction("INC"))
		//       ^^^^^^^^ invoke Dispatch in Subscribe will cause deadlock
		//store.Subscribe(func() {})
		//      ^^^^^^^^^ invoke Subscribe in Subscribe will cause panic
	})
	store.Dispatch(action.New("INC"))
}
