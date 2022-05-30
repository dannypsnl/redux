package store_test

import (
	"testing"

	"sync"

	"github.com/dannypsnl/redux/v3/store"

	"github.com/stretchr/testify/assert"
)

func counter(state int, action string) int {
	switch action {
	case "INC":
		return state + 1
	case "DEC":
		return state - 1
	default:
		return state
	}
}

type withPayload struct {
	typ     string
	payload int
}

func payload(state int, action withPayload) int {
	switch action.typ {
	case "INC":
		return state + action.payload
	case "DEC":
		return state - action.payload
	default:
		return state
	}
}

func TestStore(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		store := store.New(counter)
		if store == nil {
			t.Error("func New fail, expected it return a pointer to store instance")
		}
	})
	t.Run("Dispatch", func(t *testing.T) {
		store := store.New(counter, payload)
		store.Dispatch("INC")
		if store.StateOf(counter) != 1 {
			t.Error("counter can not work")
		}
		store.Dispatch(withPayload{
			typ:     "INC",
			payload: 10,
		})
		if store.StateOf(payload) != 10 {
			t.Error("payload should increase by payload")
		}
	})
	t.Run("StateOf", func(t *testing.T) {
		store := store.New(counter)
		store.Dispatch("DEC")
		if store.StateOf(counter) != -1 {
			t.Error("StateOf should return reducer's state")
		}
	})
	t.Run("Subscribe", func(t *testing.T) {
		store := store.New(counter)

		subcribing := make(chan bool)
		store.Subscribe(func() {
			subcribing <- true
		})
		go store.Dispatch(0)

		b := <-subcribing
		if !b {
			t.Error("Subscribe do not work")
		}
	})
	t.Run("UsingLambda", func(t *testing.T) {
		lambda := func(state int, action int) int {
			return state + action
		}
		store := store.New(lambda)
		store.Dispatch(10)
		state := store.StateOf(lambda)
		if state != 10 {
			t.Error("store can't work with lambda")
		}
	})
}

func TestConcurrencySafe(t *testing.T) {
	var wg sync.WaitGroup

	counter := func(s int, a int) int {
		return s + a
	}
	store := store.New(counter)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			store.Dispatch(1)
		}()
	}

	wg.Wait()

	actual := store.StateOf(counter)
	expected := 100

	assert.Equal(t, expected, actual)
}

func TestStoreShouldPanicWhen(t *testing.T) {
	t.Run("SubscribedFuncCallSubscribe", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("subscribed function call store.Subscribe should panic")
			}
		}()
		foo := func(s int, action int) int { return 0 }
		store := store.New(foo)
		store.Subscribe(func() {
			store.Subscribe(func() {})
		})
		store.Dispatch(0)
	})
	t.Run("HaveDuplicatedReducers", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("duplicated reducer should cause panic when New a Store")
			}
		}()
		counter := func(state int, payload int) int { return state + payload }
		store.New(counter, counter)
	})
	t.Run("ReceiveInvalidReducer", func(t *testing.T) {
		error1 := func(state int, act string) string {
			return ""
		}
		error2 := func(state string, act string) (string, error) {
			return "", nil
		}
		error3 := func() {}
		testPanic(t, error1, "should panic when reducer return a state is type different than input state type")
		testPanic(t, error2, "should panic when reducer return several type")
		testPanic(t, error3, "should panic when reducer do not contain state & action two parameters")
		testPanic(t, nil, "should panic for the invalid value")
	})
}

func testPanic(t *testing.T, reducer interface{}, msg string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(msg)
		}
	}()
	store.New(reducer)
}
