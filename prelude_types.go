package redux

import (
	"github.com/dannypsnl/redux/action"
)

type Reducer func(interface{}, action.Action) interface{}
