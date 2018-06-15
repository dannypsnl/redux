package rematch

import (
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

func NewCountingModel() *CountingModel {
	return &CountingModel{
		State: 0,
	}
}

func TestNewStoreByRematch(t *testing.T) {
	c := NewCountingModel()
	store := store.New(c)
	store.Dispatch(c.Action(c.Increase).With(10))
}
