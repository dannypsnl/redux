package redux

// Action is a type for reducer to know what should it do by recognize Action::Type
type Action struct {
	// Type help user recognize which signal they receive
	Type string
	// Args contain context of action, for example, login need account & password
	Args map[string]interface{}
}

// SendAction give a simple way to pass pointer of Action, it will be clear and performance.
func SendAction(typ string) *Action {
	return &Action{
		Type: typ,
	}
}
