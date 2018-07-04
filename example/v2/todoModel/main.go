package main

import (
	"fmt"

	"github.com/dannypsnl/redux/example/v2/todoModel/todo"

	"github.com/dannypsnl/redux/v2/store"
)

func main() {
	store := store.New(todo.Reducer)
	newTodo := todo.Reducer.Action(todo.Reducer.AddTodo)

	store.Dispatch(newTodo.With("first todo"))
	store.Dispatch(newTodo.With("second todo"))

	todos := store.StateOf(todo.Reducer).(todo.Model)

	fmt.Println("Todo List")
	for i, todo := range todos {
		if !todo.Done {
			fmt.Printf("%d. %v\n", i+1, todo.Title)
		}
	}
}
