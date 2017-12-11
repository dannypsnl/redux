package redux

import (
	"encoding/json"
	"fmt"
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
	Ext string `json:"ext"`
	Mod string `json:"mod"`
	// two mod: +x, +d
}

func fileUpdator(state interface{}, act Action) interface{} {
	if state == nil {
		return file{
			Ext: "elz",
			Mod: "+x",
		}
	}
	switch act.Type {
	case "chmod":
		if act.Args["mod"] == "+x" && !strings.Contains(state.(file).Mod, "+x") {
			return file{
				Ext: state.(file).Ext,
				Mod: state.(file).Mod + "+x",
			}
		} else if act.Args["mod"] == "+d" && !strings.Contains(state.(file).Mod, "+d") {
			return file{
				Ext: state.(file).Ext,
				Mod: state.(file).Mod + "+d",
			}
		} else if act.Args["mod"] == "-d" {
			strings.Replace(state.(file).Mod, "+d", "", -1)
		} else if act.Args["mod"] == "-x" {
			strings.Replace(state.(file).Mod, "+x", "", -1)
		} else {
			return state
		}
	case "new ext":
		if act.Args["Ext"] != nil {
			return file{
				Ext: act.Args["Ext"].(string),
				Mod: state.(file).Mod,
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
  "fileUpdator":{"ext":"elz","mod":"+x"}
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

type NestData struct {
	I int `json:"integer"`
}

type TestData struct {
	Nest NestData `json:"nest"`
}

// !!! If nest the struct, then the Marshal can't get the correct key
func TestJson(t *testing.T) {
	td := TestData{
		Nest: NestData{I: 10},
	}
	b, _ := json.Marshal(td)
	formatString := fmt.Sprintf("%s", b)
	expected := `{"nest":{"integer":10}}`
	if strings.Compare(expected, formatString) != 0 {
		t.Errorf("expected: %s, actual: %s", expected, formatString)
	}
}
