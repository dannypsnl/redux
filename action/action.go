// Package action provides Action implementation for redux
package action

// Action is a type for reducer to know what should it do by recognize Action::Type
type Action struct {
	// Type help user recognize which signal they receive
	Type string
	// Args contain context of action, for example, login need account & password
	Args map[string]interface{}
}

// New create a new action and return a pointer to it.
//
// Example:
//   action.New("Type")
func New(typ string) *Action {
	return &Action{
		Type: typ,
		Args: make(map[string]interface{}),
	}
}

// Arg help you append new Argument onto Action with fluent API
//
// Example:
//   action.New("Type").
//             Arg("author", "danny").
//             Arg("age", 20)
func (act *Action) Arg(arg string, val interface{}) *Action {
	act.Args[arg] = val
	return act
}
