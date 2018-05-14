package rematch

import (
	"fmt"

	"github.com/dannypsnl/redux/v1"
	"github.com/dannypsnl/redux/v1/action"
)

// Reducer provide a readable structure help you build reducer more rematchable
//
// Example:
//   var r rematch.Reducer
//   func init() {
//       r = Reducer {
//           State: 0,
//           Reducers: rematch.Reducers {
//               "INC": func(s interface{}, a action.Action) interface{} {
//                   return s.(int) + 1
//               }
//           },
//       }
//   }
//
// !Note: the limit at here is store using function address generate reducer name, so we have to use Reduce and a func proxy to using rematch Reducer
type Reducer struct {
	State    interface{}
	Reducers Reducers
}

// Reduce return store acceptable reducer(func form)
//
// Usage:
//   func reducer(s interface{}, a action.Action) interface{} {
//       return r.Reduce()(s, a)
//   }
func (r *Reducer) Reduce() v1.Reducer {
	return func(state interface{}, act action.Action) interface{} {
		if state == nil {
			return r.State
		}
		r, ok := r.Reducers[act.Type]
		if ok {
			return r(state, act)
		}
		return state
	}
}

// Action in rematch Reducer help us detect only the action exist in this rematch Reducer can be using
//
// Else will panic
func (r *Reducer) Action(typ string) *action.Action {
	for k, _ := range r.Reducers {
		if k == typ {
			return action.New(typ)
		}
	}
	panic(fmt.Sprintf("Action %s is not the legal action for this rematch reducer", typ))
}

// Reducers let we can meta reducer group much more readable
type Reducers map[string]v1.Reducer
