package store

import (
	"github.com/dannypsnl/assert"
	"testing"

	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/rematch"
)

var reCounter rematch.Reducer

func init() {
	reCounter = rematch.Reducer{
		State: 0,
		Reducers: rematch.Reducers{
			"INC": func(state interface{}, act action.Action) interface{} {
				return state.(int) + act.Args["payload"].(int)
			},
			"DEC": func(state interface{}, act action.Action) interface{} {
				return state.(int) - act.Args["payload"].(int)
			},
		},
	}
}

func TestWorkWithV1Rematch(t *testing.T) {
	assert := assert.NewTester(t)

	wrap := func(s interface{}, act action.Action) interface{} {
		return reCounter.Reduce()(s, act)
	}
	store := New(wrap)
	store.Dispatch(*reCounter.
		Action("INC").
		Arg("payload", 10))

	assert.Eq(store.StateOf(wrap), 10)
}
