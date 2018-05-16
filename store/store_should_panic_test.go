package store

import (
	"testing"

	"github.com/dannypsnl/redux/action"
)

func TestDuplicatedReducerShouldCausePanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`duplicated reducer should break the process`)
		}
	}()
	store := /*store.*/ New(counter, counter)
	store.Dispatch(action.New("INC"))
}

func TestCallSubscribeInSubscribtorShouldPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`should panic when subscribetor trying to call store::Subscribe`)
		}
	}()

	store := /*store.*/ New(counter)
	store.Subscribe(func() {
		store.Subscribe(func() {})
	})

	store.Dispatch(action.New("INC"))
}
