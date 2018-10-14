package rematch

import (
	"reflect"
	"runtime"
	"strings"
)

// Reducer is core of rematch.
//
// Usage:
//  type CountingModel struct {
//      rematch.Reducer
//      State int
//  }
//
//  func (cm *CountingModel) Increase(s int, payload int) int {
//      return s + payload
//  }
//
//  // main
//  c := &CountingModel{
//      State: 0,
//  }
//  store := store.New(c)
//  store.Dispatch(c.Action(c.Increase).With(10))
//  store.StateOf(c) // should be 10
type Reducer struct {
	ms map[string]reflect.Value
}

func (r Reducer) methods(v interface{}) map[string]reflect.Value {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	methods := make(map[string]reflect.Value)
	for i := 1; i < rt.NumMethod(); i++ {
		m := rt.Method(i) // rt.Method.Func return func with first argument as receiver
		mt := m.Type
		if mt.NumIn() == 3 &&
			mt.NumOut() == 1 &&
			mt.In(1) == mt.Out(0) {
			// rv.Method return func with now receiver
			typN := rt.String()[1:]
			locOfDot := strings.LastIndex(typN, ".")
			r := typN[:locOfDot+1] + "(*" + typN[locOfDot+1:] + ")"
			methods[r+"."+m.Name] = rv.Method(i)
		}
	}
	return methods
}

// actions would find out all possible action in user's rematcher & initial it
func (r Reducer) actions(v interface{}) {
	rv := reflect.ValueOf(v).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		value, ok := rt.Field(i).Tag.Lookup("action")
		if ok {
			if rv.Field(i).Type() != reflect.TypeOf(&Action{}) {
				panic("type of field with action tag must be `*rematch.Action`")
			}
			fullPath := rt.PkgPath()
			pkg := fullPath[strings.LastIndexByte(fullPath, '/')+1:]
			funcName := pkg + ".(*" + rt.Name() + ")." + value
			rv.Field(i).Set(reflect.ValueOf(&Action{
				funcName: funcName,
			}))
		}
	}
}

// InsideReducer is not preparing for you, it's because reflect can't see private method, so export it
func (r Reducer) InsideReducer(v interface{}) func(interface{}, *Action) interface{} {
	r.actions(v)
	r.ms = r.methods(v)
	return func(state interface{}, action *Action) interface{} {
		if m, ok := r.ms[action.reducerName()]; ok {
			return m.Call(
				[]reflect.Value{
					reflect.ValueOf(state),
					reflect.ValueOf(action.payload()),
				},
			)[0].Interface()
		}
		return state
	}
}

// Action return a new `rematch.Action` with the method will be executed(as type)
//
// Then you can use `Action.With` to creating payload
func (r Reducer) Action(method interface{}) *Action {
	return &Action{
		funcName: getReducerName(method),
	}
}

// Action type is rematch dispatcher, it store key to find correct internal reducer & payload for it
type Action struct {
	funcName string
	with     interface{}
}

// With help you insert any payload you want into Action.
// Just remind payload should has the same type with internal reducer expected
//
// For example:
//
//  type User struct {
//      Reducer
//      State string
//  }
//
//  func (m *Model) Rename(s string, newName string) string {
//      return newName
//  }
//
// You should `With` a `string` rather than put a `int`
//
// Usage:
//   store.Dispatch(c.Action(c.Increase).With(10))
func (a *Action) With(payload interface{}) *Action {
	a.with = payload
	return a
}

func (a Action) reducerName() string {
	return a.funcName
}

func (a Action) payload() interface{} {
	return a.with
}

// getReducerName is a helper func to get function's ref name.
func getReducerName(r interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
	// fullName's format is `path/to/package.function_name`
	// we don't want path/to part.
	// package is full path(GOPATH/src/package) to it
	// len-3 is because a method contains suffix `-fm`
	return fullName[strings.LastIndexByte(fullName, '/')+1 : len(fullName)-3]
}
