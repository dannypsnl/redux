package store

import (
	"github.com/dannypsnl/assert"
	"testing"

	"github.com/dannypsnl/redux/v2/rematch"
)

type user struct {
	name string
	age  int
}

type UserModel struct {
	rematch.Reducer
	State user
}

func (cm *UserModel) UpdateAge(state user, newAge int) user {
	return user{
		name: state.name,
		age:  newAge,
	}
}

func (cm *UserModel) UpdateName(state user, newName string) user {
	return user{
		name: newName,
		age:  state.age,
	}
}

func NewCountingModel() *UserModel {
	return &UserModel{
		State: user{
			name: "danny",
			age:  20,
		},
	}
}

func TestNewStoreByRematch(t *testing.T) {
	assert := assert.NewTester(t)

	c := NewCountingModel()
	store := New(c)
	// yp, I'm getting older, lalala
	store.Dispatch(c.Action(c.UpdateAge).With(21))
	store.Dispatch(c.Action(c.UpdateName).With("Danny"))

	actual := store.StateOf(c)
	expect := user{
		name: "Danny",
		age:  21,
	}
	assert.Eq(actual, expect)
}

func TestNewWithDuplicateRematcher(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect duplicated rematch.Reducer should cause panic but didn't.")
		}
	}()
	c := NewCountingModel()
	New(c, c)
}

func TestNewWithNotARematcher(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("New store with illegl reducer should panic but didn't.")
		}
	}()
	c := &user{}
	New(c)
}
