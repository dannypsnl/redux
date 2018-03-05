// Package store provides the store implement for redux.
package store

import (
	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/middleware"
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
	// disMu Lock each Dispatch call
	disMu sync.Mutex
	// subMu Lock each Subscribe call
	subMu sync.Mutex
	// invokeSubscribeInSubscribtor log should process panic or not, because we have to help test cache it
	invokeSubscribeInSubscribtor bool
	// onDispatching make sure Subscribetor can't call Subscribe
	onDispatching bool
	middleware    middleware.Next
}

// New create a Store by reducers
//
// Usage:
//   store := store.New(reducer...)
func New(r reducer, reducers ...reducer) *Store {
	s := &Store{
		state:      make(map[string]interface{}),
		middleware: func(a *action.Action) *action.Action { return a },
	}
	s.emit(r)
	for _, r := range reducers {
		s.emit(r)
	}
	return s
}

// Inter comment
// emit append r into Store's reducers
func (s *Store) emit(r reducer) {
	// initial state will be return when current state is nil, so we send nil at here.
	reducerState := s.state[getReducerName(r)]
	if reducerState != nil {
		panic("You can not create duplicated reducer")
	}
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
	// Make sure dispatch process won't be terminated
	s.disMu.Lock()
	defer s.disMu.Unlock()

	s.updateState(act)
	s.onDispatching = true

	var wg sync.WaitGroup
	// we call subscribed function after state updated.
	wg.Add(len(s.subscribes))
	for _, subscribtor := range s.subscribes {
		go func(su func()) {
			defer func() {
				if p := recover(); p != nil {
					s.invokeSubscribeInSubscribtor = true
				}
				wg.Done()
			}()
			su()
		}(subscribtor)
	}
	wg.Wait()
	// Have to check at here then panic is because recover can't recover second level's panic
	// We need to test this part's code, so have recover then panic at here
	if s.invokeSubscribeInSubscribtor {
		panic(`You may not call store.subscribe() while the reducer is executing!`)
	}
	s.onDispatching = false
}

func (s *Store) updateState(act *action.Action) {
	// dispatch action to every reducer, then reducer update it's mapping state.
	for _, r := range s.reducers {
		funcName := getReducerName(r) // getReducerName in util.go
		nowState := s.state[funcName]
		a := s.middleware(act)
		s.state[funcName] = r(nowState, *a)
	}
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
	if s.onDispatching {
		panic(`You may not call store.subscribe() while the reducer is executing.`)
	}
	s.subscribes = append(s.subscribes, subscribetor)
}

// middlewareType is real middleware shape
type middlewareType func(*Store) middleware.Middleware

// ApplyMiddleware expected middlewares and use it update store's middleware to control action
func (s *Store) ApplyMiddleware(middlewares ...middlewareType) {
	for _, middleware := range middlewares {
		s.middleware = middleware(s)(s.middleware)
	}
}
