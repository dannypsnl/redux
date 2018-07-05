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
			methods[rt.String()[1:]+"."+m.Name] = rv.Method(i)
		}
	}
	return methods
}

// InsideReducer is not preparing for you, it's because reflect can't see private method, so export it
func (r Reducer) InsideReducer(v interface{}) func(interface{}, *action) interface{} {
	r.ms = r.methods(v)
	return func(state interface{}, action *action) interface{} {
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

// Action return a new `rematch.action` with the method will be executed(as type)
//
// Then you can use `action.With` to creating payload
func (r Reducer) Action(method interface{}) *action {
	return &action{
		funcName: getReducerName(method),
	}
}

type action struct {
	funcName string
	with     interface{}
}

// With help you insert any payload you want into action.
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
//   store.Dispatch(c.action(c.Increase).With(10))
func (a *action) With(payload interface{}) *action {
	a.with = payload
	return a
}

func (a action) reducerName() string {
	return a.funcName
}

func (a action) payload() interface{} {
	return a.with
}

// getReducerName is a helper func to get function's ref name.
func getReducerName(r interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
	pkgToMethod := fullName[strings.LastIndexByte(fullName, '/')+1 : len(fullName)-3]
	star := strings.IndexRune(pkgToMethod, '*')
	leftParent := strings.IndexRune(pkgToMethod, '(')
	rightParent := strings.IndexRune(pkgToMethod, ')')
	// fullName's format is `package.function_name`
	// we don't want package part.
	// package is full path(GOPATH/src/package_part) to it
	// len-3 is because a method contains suffix `-fm`
	return pkgToMethod[:leftParent] + pkgToMethod[star+1:rightParent] + pkgToMethod[rightParent+1:]
}
