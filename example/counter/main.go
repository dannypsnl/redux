package main

import (
	"fmt"

	"github.com/dannypsnl/redux/v3/store"
)

func main() {
	counter := func(state, payload int) int {
		return state + payload
	}
	store := store.New(counter)
	store.Dispatch(100)
	store.Dispatch(-50)
	fmt.Printf("%d\n", store.StateOf(counter)) // should print out: 50
}
