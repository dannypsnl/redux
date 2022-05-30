package store

import (
	"reflect"
	"sync"
)

type Store[State, Action any] struct {
	reducers            map[uintptr]func(State, Action) State
	state               map[uintptr]State
	subscribedFuncs     []func()
	dispatch            sync.Mutex
	subscribe           sync.Mutex
	subscribedFuncPanic bool
	onDispatching       bool
}

// New create a Store by reducers
func New[State, Action any](reducers ...func(State, Action) State) *Store[State, Action] {
	newStore := &Store[State, Action]{
		reducers: make(map[uintptr]func(State, Action) State),
		state:    make(map[uintptr]State),
	}
	for _, reducer := range reducers {
		r := reflect.ValueOf(reducer)
		if _, ok := newStore.state[r.Pointer()]; ok {
			panic("You can't put duplicated reducer into the same store!")
		}

		var initState State
		newStore.reducers[r.Pointer()] = reducer
		newStore.state[r.Pointer()] = initState
	}
	return newStore
}

// Dispatch send action to reducers in Store to update state
//
// In v3, action can be anything you want
//
// For example:
//   func counter(state, payload int) int {
//   	return state + payload
//   }
//
// Action type should be `int`, if you use others type, just use others type
func (s *Store[State, Action]) Dispatch(action Action) {
	s.dispatch.Lock()
	defer s.dispatch.Unlock()

	for at, reducer := range s.reducers {
		// action type matches second argument
		reducer(s.state[at], action)
	}

	s.onDispatching = true

	var wg sync.WaitGroup
	wg.Add(len(s.subscribedFuncs))
	for _, subscribedFunc := range s.subscribedFuncs {
		go func(f func()) {
			defer func() {
				if r := recover(); r != nil {
					s.subscribedFuncPanic = true
				}
				wg.Done()
			}()
			f()
		}(subscribedFunc)
	}
	wg.Wait()

	if s.subscribedFuncPanic {
		panic("You can't call Subscribe in subscribed function")
	}
	s.onDispatching = false
}

// Subscribe emit a function will be triggered after executing Dispatch
func (s *Store[State, Action]) Subscribe(function func()) {
	s.subscribe.Lock()
	defer s.subscribe.Unlock()
	if s.onDispatching {
		panic("You can't call Subscribe in subscribed function")
	}
	s.subscribedFuncs = append(s.subscribedFuncs, function)
}

// StateOf return the reducer matches state
func (s *Store[State, Action]) StateOf(reducer interface{}) State {
	ofReducer := reflect.ValueOf(reducer).Pointer()
	return s.state[ofReducer]
}
