package redux

import "testing"

var users = map[string]string{
	"danny": "1234",
}

func login(state interface{}, action Action) interface{} {
	if state == nil {
		return "Guest"
	}
	switch action.Type {
	case "login":
		if users[action.Args["user"].(string)] == action.Args["password"] {
			return action.Args["user"].(string) + " Login"
		}
		return state
	default:
		return state
	}
}

func TestSerialize(t *testing.T) {
	store := NewStore(counter, login)
	expected := "{\n  \"counter\":0,\n  \"login\":\"Guest\"\n}"
	if store.JSON() != expected {
		t.Errorf("serialized result is not expected, expected: %s, actual: %s", expected, store.JSON())
	}
	store.Dispatch(SendAction("INC"))
	store.Dispatch(SendAction("INC"))
	store.Dispatch(&Action{
		Type: "login",
		Args: map[string]interface{}{
			"user":     "danny",
			"password": "1234",
		},
	})
	expected = "{\n  \"counter\":2,\n  \"login\":\"danny Login\"\n}"
	if store.JSON() != expected {
		t.Errorf("serialized result is not expected, expected: %s, actual: %s", expected, store.JSON())
	}
}
