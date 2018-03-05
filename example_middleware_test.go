package redux_test

import (
	"fmt"

	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/middleware"
	"github.com/dannypsnl/redux/store"
)

func logger(store *store.Store) middleware.Middleware {
	return func(next middleware.Next) middleware.Next {
		return func(act *action.Action) *action.Action {
			fmt.Printf("Dispatching %v", *act)
			r := next(act)
			return r
		}
	}
}

func Example_middleware() {
	store := store.New(counter).
		ApplyMiddleware(logger)
	store.Dispatch(action.New("INC"))
	// Output:
	// Dispatching {INC map[]}
}
