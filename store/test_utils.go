package store

import (
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

func jump(state interface{}, action action.Action) interface{} {
	if state == nil {
		return "TOP"
	}
	switch action.Type {
	case "JUMP":
		return "TOP"
	case "FALL":
		return "DOWN"
	default:
		return state
	}
}

func foo(state interface{}, act action.Action) interface{} {
	if state == nil {
		return func() {}
	}
	return state
}
