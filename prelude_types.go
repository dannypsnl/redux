package redux

import (
	"github.com/dannypsnl/redux/v1/action"
)

type Reducer func(interface{}, action.Action) interface{}
