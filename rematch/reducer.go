package rematch

import (
	"fmt"

	"github.com/dannypsnl/redux"
	"github.com/dannypsnl/redux/action"
)

type Reducer struct {
	State    interface{}
	Reducers Reducers
}

func (r *Reducer) Action(typ string) *action.Action {
	for k, _ := range r.Reducers {
		if k == typ {
			return action.New(typ)
		}
	}
	panic(fmt.Sprintf("Action %s is not the legal action for this rematch reducer", typ))
}

type Reducers map[string]redux.Reducer
