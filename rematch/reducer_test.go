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
	return c.Reduce()(s, a)
}

func testWorkWithStore(t *testing.T) {
	store := store.New(counter)

	testInitStore(t, store)
	testDispatch(t, store, c.Action("INC").Arg("payload", 10), 10)
	testDispatch(t, store, action.New("PLUS"), 10)
}

func testInitStore(t *testing.T, store *store.Store) {
	if store.GetState("counter") != 0 {
		t.Error("store::New can't work with rematch::Reducer")
	}
}
func testDispatch(t *testing.T, store *store.Store, act *action.Action, expected interface{}) {
	store.Dispatch(act)
	if store.GetState("counter") != expected {
		t.Error("store::Dispatch can't work with rematch::Reducer")
	}
}
