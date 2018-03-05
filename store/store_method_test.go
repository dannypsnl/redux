package store

import (
	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/middleware"
	"sync"
	"testing"
)

func TestDispatchInConcurrencyIsSafe(t *testing.T) {
	store := /*store.*/ New(counter)
	n := 1000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			store.Dispatch(action.New("INC"))
		}()
	}
	wg.Wait()
	if store.GetState("counter") != 1000 {
		t.Errorf("Expected: %d, Actual: %d", 1000, store.GetState("counter"))
	}
}

func TestStoreStateCanBeUpdateByDispatch(t *testing.T) {
	thisState := "jump"
	expectedState := "TOP"

	store := /*store.*/ New(counter, jump)
	store.Dispatch(action.New("JUMP"))

	if store.GetState(thisState) != expectedState {
		t.Errorf("Expected: %v, Actual: %v", expectedState, store.GetState(thisState))
	}
}

func TestFall(t *testing.T) {
	store := /*store.*/ New(jump)
	store.Dispatch(action.New("FALL"))

	if store.GetState("jump") != "DOWN" {
		t.Error("Wow! You can fly")
	}
}

func increaseToDec(store *Store) middleware.Middleware {
	return func(next middleware.Next) middleware.Next {
		return func(act *action.Action) *action.Action {
			switch act.Type {
			case "INC":
				return next(action.New("DEC"))
			default:
				return next(action.New("INC"))
			}
		}
	}
}

func TestMiddlewareFirstTry(t *testing.T) {
	store := /*store.*/ New(counter).
		ApplyMiddleware(increaseToDec)
	store.Dispatch(action.New("INC"))

	if store.GetState("counter") != -1 {
		t.Errorf("error, actual: %v", store.GetState("counter"))
	}
}
