package redux

// Action is a type for reducer to know what should it do by recognize Action::Type
type Action struct {
	Type string
	Args map[string]interface{}
}

// SendAction give a simple way to pass pointer of Action, it will be clear and performance.
func SendAction(typ string) *Action {
	return &Action{
		Type: typ,
	}
}

// reducer is a function get current state and a Action, return new state.
type reducer func(interface{}, Action) interface{}
