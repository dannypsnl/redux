package user

import (
	"github.com/dannypsnl/redux/v2/rematch"
)

type Model struct {
	Name string
	Age  int
}

type UserModel struct {
	rematch.Reducer
	State Model
}

func (cm *UserModel) UpdateName(s Model, newName string) Model {
	return Model{
		Name: newName,
		Age:  s.Age,
	}
}

func (cm *UserModel) UpdateAge(s Model, newAge int) Model {
	return Model{
		Name: s.Name,
		Age:  newAge,
	}
}

var Reducer *UserModel

func init() {
	Reducer = &UserModel{
		State: Model{
			Name: "dan",
			Age:  20,
		},
	}
}
