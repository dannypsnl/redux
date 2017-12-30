// Package store provides the store implement for redux.
package store

import (
	"github.com/dannypsnl/redux/action"
	"sync"
)

// Inter comment
// reducer is a function get current state and a Action, return new state.
type reducer func(interface{}, action.Action) interface{}

// Store is a type manage your data
type Store struct {
	// state contain key map state
	// and we use key to spread reducer's state
	// for example: "counter" mapping to reducer named counter.
	// when use GetState("counter"), we will got the current state of "counter" key's mapping target.
	state map[string]interface{}
	// reducers contains every reducer of this Store.
	reducers []reducer
	// subscribes contains those function we want to invoke at dispatch
	subscribes []func()
	// mu Lock each Dispatch call
	mu sync.Mutex
	// subMu Lock each Subscribe call
	subMu sync.Mutex
	// shouldPanic log should process panic or not, because we have to help test cache it
	shouldPanic bool
	// isDispatching make sure Subscribetor can't call Subscribe
	isDispatching bool
}

// New create a Store by reducers
//
// Usage:
//   store := store.New(reducer...)
func New(r reducer, reducers ...reducer) *Store {
	s := &Store{
		state: make(map[string]interface{}),
	}
	s.newReducer(r)
	for _, r := range reducers {
		s.newReducer(r)
	}
	return s
}

// Inter comment
// newReducer append r into Store's reducers
func (s *Store) newReducer(r reducer) {
	// initial state will be return when current state is nil, so we send nil at here.
	s.state[getReducerName(r)] = r(nil, action.Action{})
	s.reducers = append(s.reducers, r)
}

// GetState recieve string key to get the mapped state
//
// Usage:
//  store.GetState("reducer_name")
func (s *Store) GetState(key string) interface{} {
	return s.state[key]
}

// Dispatch send action to every reducer
//
// Usage:
//  store.Dispatch(action.New("Type").
//                        Arg("key", expression))
func (s *Store) Dispatch(act *action.Action) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// we dispatch action to every reducer, and reducer update mapping state.
	for _, r := range s.reducers {
		funcName := getReducerName(r) // getReducerName in util.go
		s.state[funcName] = r(s.state[funcName], *act)
	}
	s.isDispatching = true
	var wg sync.WaitGroup
	// we call subscribed function after state updated.
	for _, subscribtor := range s.subscribes {
		wg.Add(1)
		go func(su func()) {
			defer func() {
				if p := recover(); p != nil {
					s.shouldPanic = true
				}
				wg.Done()
			}()
			su()
		}(subscribtor)
	}
	wg.Wait()
	if s.shouldPanic {
		panic(`You may not call store.subscribe() while the reducer is executing!`)
	}
	s.isDispatching = false
}

// Subscribe emit argument into subscribes chain, they will be invoked in Dispatch.
//
// Usage:
//  store.Subscribe(func() {
//      // Do something...
//  })
//
// !Warning, subscribed function can't invoke Dispatch, it will deadlock
//
// !Warning, subscribed function can't invoke Subscribe, it will panic
func (s *Store) Subscribe(subscribetor func()) {
	s.subMu.Lock()
	defer s.subMu.Unlock()
	if s.isDispatching {
		panic(`You may not call store.subscribe() while the reducer is executing.`)
	}
	s.subscribes = append(s.subscribes, subscribetor)
}
