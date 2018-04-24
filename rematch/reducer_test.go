package rematch

import (
	"testing"

	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/store"
)

var c Reducer

func init() {
	c = Reducer{
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
}

func TestRematchReducer(t *testing.T) {
	testInit(t, c)
	testAction(t, c)
	testNotExistAction(t, c)
	testWorkWithStore(t)
}

func testInit(t *testing.T, counter Reducer) {
	if counter.State != 0 {
		t.Error("initialize bug")
	}
}
func testAction(t *testing.T, counter Reducer) {
	act := counter.Action("INC").Arg("payload", 10)
	if act.Type != "INC" || act.Args["payload"] != 10 {
		t.Error("rematch::Reducer can't generate correct action")
	}
}
func testNotExistAction(t *testing.T, counter Reducer) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("generate the action not existing in rematch::Reducer should cause panic")
		}
	}()
	counter.Action("PLUS")
}

func counter(s interface{}, a action.Action) interface{} {
	return c.Reducer()(s, a)
}

func testWorkWithStore(t *testing.T) {
	store := store.New(counter)
	if store.GetState("counter") != 0 {
		t.Error("store can't work with rematch reducer")
	}
}
