package main

import (
	"fmt"

	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

type CountingModel struct {
	rematch.Reducer
	State int
}

func (cm *CountingModel) Increase(s, payload int) int {
	return s + payload
}

func (cm *CountingModel) Decrease(s, payload int) int {
	return s - payload
}

func main() {
	c := &CountingModel{
		State: 0,
	}
	store := store.New(c)
	store.Dispatch(c.Action(c.Increase).With(10))
	store.Dispatch(c.Action(c.Increase).With(10))
	store.Dispatch(c.Action(c.Decrease).With(30))

	fmt.Printf("result: %d\n", store.StateOf(c)) // expect: -10
}
