package user

import (
	"github.com/dannypsnl/redux/v3/rematch"
)

type Model struct {
	Name string
	Age  int
}

type userModel struct {
	rematch.Reducer
	State Model
}

func (cm *userModel) UpdateName(s Model, newName string) Model {
	return Model{
		Name: newName,
		Age:  s.Age,
	}
}

func (cm *userModel) UpdateAge(s Model, newAge int) Model {
	return Model{
		Name: s.Name,
		Age:  newAge,
	}
}

var Reducer *userModel

func init() {
	Reducer = &userModel{
		State: Model{
			Name: "dan",
			Age:  20,
		},
	}
}
