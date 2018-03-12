package store

import (
	"github.com/dannypsnl/redux/action"
	"strings"
)

func counter(state interface{}, action action.Action) interface{} {
	if state == nil {
		return 0
	}
	switch action.Type {
	case "INC":
		return state.(int) + 1
	case "DEC":
		return state.(int) - 1
	default:
		return state
	}
}

func jump(state interface{}, action action.Action) interface{} {
	if state == nil {
		return "TOP"
	}
	switch action.Type {
	case "JUMP":
		return "TOP"
	case "FALL":
		return "DOWN"
	default:
		return state
	}
}

var users = map[string]string{
	"danny": "1234",
}

func login(state interface{}, action action.Action) interface{} {
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

type file struct {
	// Ext mean extension
	Ext string `json:"ext"`
	// two mod: +x, +d
	Mod string `json:"mod"`
}

func fileUpdator(state interface{}, act action.Action) interface{} {
	if state == nil {
		return &file{
			Ext: "elz",
			Mod: "+x",
		}
	}
	switch act.Type {
	case "chmod":
		if act.Args["mod"] == "+x" && !strings.Contains(state.(file).Mod, "+x") {
			return &file{
				Ext: state.(file).Ext,
				Mod: state.(file).Mod + "+x",
			}
		} else if act.Args["mod"] == "+d" && !strings.Contains(state.(file).Mod, "+d") {
			return &file{
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
			return &file{
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
