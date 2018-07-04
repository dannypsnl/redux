package todo

import (
	"github.com/dannypsnl/redux/v2/rematch"
)

var Reducer *todoModel

func init() {
	Reducer = &todoModel{
		State: make([]Todo, 0),
	}
}

type Todo struct {
	Title string
	Done  bool
}

type Model []Todo

type todoModel struct {
	rematch.Reducer
	State Model
}

func (todo *todoModel) AddTodo(state Model, title string) Model {
	return append(state, Todo{Title: title})
}
