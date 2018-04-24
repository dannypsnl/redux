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

func (r *Reducer) Action(typ string) *action.Action {
	return action.New(typ)
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
func testAction(t *testing.T, counter Reducer) {
	act := counter.Action("INC").Arg("payload", 10)
	if act.Type != "INC" || act.Args["payload"] != 10 {
		t.Error()
	}
}
