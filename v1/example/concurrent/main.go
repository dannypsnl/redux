package main

import (
	"fmt"
	"sync"

	"github.com/dannypsnl/redux/v1/action"
	"github.com/dannypsnl/redux/v1/store"
)

func counter(state interface{}, action action.Action) interface{} {
	// initial state
	if state == nil {
		return 0
	}
	switch action.Type {
	case "INC":
		return state.(int) + 1
	case "DEC":
		return state.(int) - 1
	default:
		return state
	}
}

func main() {
	store := store.New(counter)
	store.Subscribe(func() {
		fmt.Printf("State is %v\n", store.GetState("counter"))
	})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			store.Dispatch(action.New("INC"))
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			store.Dispatch(action.New("DEC"))
			wg.Done()
		}()
	}
	wg.Wait()
}
