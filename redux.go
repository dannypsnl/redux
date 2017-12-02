package redux

import (
	"sync"
)

type Action struct {
	Type string
}

// SendAction give a simple way to pass pointer of Action, it will be clear and performance.
func SendAction(typ string) *Action {
	return &Action{
		Type: typ,
	}
}

// Reducer is a function get current state and a Action, return new state.
type Reducer func(interface{}, Action) interface{}

// store is a type manage our data
type store struct {
	// reducers contains every reducer of this store.
	reducers []Reducer
	// GetState contain key map state
	// and we use key to spread reducer's state
	// for example: "counter" mapping to reducer named counter.
	// when use GetState["counter"], we will got the current state of "counter" key's mapping target.
	GetState map[string]interface{}
	// subscribes contains those function we want to invoke at dispatch
	subscribes  []func()
	atSubscribe bool
	mu          sync.Mutex
}

func NewStore(reducer Reducer, reducers ...Reducer) *store {
	s := &store{
		GetState:    make(map[string]interface{}),
		atSubscribe: false,
	}
	s.newReducer(reducer)
	for _, reducer := range reducers {
		s.newReducer(reducer)
	}
	return s
}

func (s *store) newReducer(reducer Reducer) {
	// initial state will be return when current state is nil, so we send nil at here.
	s.GetState[getReducerName(reducer)] = reducer(nil, Action{})
	s.reducers = append(s.reducers, reducer)
}

func (s *store) Dispatch(act *Action) {
	s.mu.Lock()
	if s.atSubscribe {
		panic(`you're trying to invoke Dispatch inside the subscribed function`)
	}
	// we dispatch action to every reducer, and reducer update mapping state.
	for _, reducer := range s.reducers {
		func_name := getReducerName(reducer)
		s.GetState[func_name] = reducer(s.GetState[func_name], *act)
	}
	// we call subscribed function after state updated.
	s.atSubscribe = true
	for _, subscribtor := range s.subscribes {
		subscribtor()
	}
	s.atSubscribe = false
	s.mu.Unlock()
}

func (s *store) Subscribe(subscribe_fn func()) {
	s.subscribes = append(s.subscribes, subscribe_fn)
}
