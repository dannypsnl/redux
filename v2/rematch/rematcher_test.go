package rematch

import (
	"github.com/dannypsnl/assert"
	"testing"

	"github.com/dannypsnl/redux/v2/store"
)

type CountingModel struct {
	Reducer
	State int
}

func (cm *CountingModel) Increase(state int, payload int) int {
	return state + payload
}

func (cm *CountingModel) Decrease(state int, payload int) int {
	return state - payload
}

func NewCountingModel() *CountingModel {
	return &CountingModel{
		State: 0,
	}
}

func TestNewStoreByRematch(t *testing.T) {
	assert := assert.NewTester(t)

	c := NewCountingModel()
	store := store.New(c)
	store.Dispatch(c.Action(c.Increase).With(10))
	store.Dispatch(c.Action(c.Decrease).With(5))

	actual := store.StateOf(c)
	expect := 5
	assert.Eq(actual, expect)
}

type duplicateMethod1 struct {
	Reducer
	State int
}

type duplicateMethod2 struct {
	Reducer
	State int
}

func (d1 *duplicateMethod1) Foo(state int, payload int) int {
	return state + payload
}
func (d2 *duplicateMethod2) Foo(state int, payload int) int {
	return state + payload
}

func TestDuplicatedMethodNameWontCauseBothStatesUpdated(t *testing.T) {
	assert := assert.NewTester(t)

	d1 := &duplicateMethod1{State: 0}
	d2 := &duplicateMethod2{State: 0}

	store := store.New(d1, d2)
	store.Dispatch(d1.Action(d1.Foo).With(10))

	assert.Eq(store.StateOf(d1), 10)
	assert.Eq(store.StateOf(d2), 0)
}
