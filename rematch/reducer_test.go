package rematch

import (
	"testing"

	"github.com/dannypsnl/redux"
	"github.com/dannypsnl/redux/action"
)

type Reducer struct {
	State    interface{}
	Reducers Reducers
}
type Reducers map[string]redux.Reducer

func TestRematchReducer(t *testing.T) {
	counter := Reducer{
		State: 0,
		Reducers: Reducers{
			"INC": func(state interface{}, act action.Action) interface{} {
				return state.(int) + act.Args["payload"].(int)
			},
			"DEC": func(state interface{}, act action.Action) interface{} {
				return state.(int) - act.Args["payload"].(int)
			},
		},
	}
	testInit(t, counter)
}

func testInit(t *testing.T, counter Reducer) {
	if counter.State != 0 {
		t.Error("initialize bug")
	}
}
