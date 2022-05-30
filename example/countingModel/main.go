package main

import (
	"fmt"

	"github.com/dannypsnl/redux/v3/rematch"
	"github.com/dannypsnl/redux/v3/store"
)

type CountingModel struct {
	rematch.Reducer
	State int

	Increase *rematch.Action `action:"IncreaseImpl"`
	Decrease *rematch.Action `action:"DecreaseImpl"`
}

func (cm *CountingModel) IncreaseImpl(s, payload int) int {
	return s + payload
}

func (cm *CountingModel) DecreaseImpl(s, payload int) int {
	return s - payload
}

func main() {
	c := &CountingModel{
		State: 0,
	}
	store := store.New(c)
	store.Dispatch(c.Increase.With(10))
	store.Dispatch(c.Increase.With(10))
	store.Dispatch(c.Decrease.With(30))

	fmt.Printf("result: %d\n", store.StateOf(c)) // expect: -10
}
