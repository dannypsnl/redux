package redux

import (
	_ "fmt"
	"strings"
	"testing"
)

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
	if strings.Compare(store.JSON(), expected) != 0 {
		t.Errorf("serialized result is not expected, expected:\n`%s`, actual:\n`%s`", expected, store.JSON())
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
	if strings.Compare(store.JSON(), expected) != 0 {
		t.Errorf("serialized result is not expected, expected:\n`%s`, actual:\n`%s`", expected, store.JSON())
	}
}

type file struct {
	ext string
	mod string
	// two mod: +x, +d
}

func fileUpdator(state interface{}, act Action) interface{} {
	if state == nil {
		return file{
			ext: "elz",
			mod: "+x",
		}
	}
	switch act.Type {
	case "chmod":
		if act.Args["mod"] == "+x" && !strings.Contains(state.(file).mod, "+x") {
			return file{
				ext: state.(file).ext,
				mod: state.(file).mod + "+x",
			}
		} else if act.Args["mod"] == "+d" && !strings.Contains(state.(file).mod, "+d") {
			return file{
				ext: state.(file).ext,
				mod: state.(file).mod + "+d",
			}
		} else if act.Args["mod"] == "-d" {
			strings.Replace(state.(file).mod, "+d", "", -1)
		} else if act.Args["mod"] == "-x" {
			strings.Replace(state.(file).mod, "+x", "", -1)
		} else {
			return state
		}
	case "new ext":
		if act.Args["ext"] != nil {
			return file{
				ext: act.Args["ext"].(string),
				mod: state.(file).mod,
			}
		}
		return state
	default:
		return state
	}
	return state
}

func TestSerializeStruct(t *testing.T) {
	store := NewStore(fileUpdator)
	expected := `{
  "fileUpdator":{ext:elz mod:+x}
}`
	if strings.Compare(store.JSON(), expected) != 0 {
		t.Errorf("expected: %s, actual: %s", expected, store.JSON())
	}
	store.Dispatch(&Action{
		Type: "new ext",
		Args: map[string]interface{}{
			"ext": "cpp",
		},
	})
}
