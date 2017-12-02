package redux

import (
	"sync"
)

// Action is a type for reducer to know what should it do by recognize Action::Type
type Action struct {
	Type string
}

// SendAction give a simple way to pass pointer of Action, it will be clear and performance.
func SendAction(typ string) *Action {
	return &Action{
		Type: typ,
	}
}

// reducer is a function get current state and a Action, return new state.
type reducer func(interface{}, Action) interface{}

// store is a type manage our data
type store struct {
	// reducers contains every reducer of this store.
	reducers []reducer
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

// NewStore create a store by reducers
func NewStore(r reducer, reducers ...reducer) *store {
	s := &store{
		GetState:    make(map[string]interface{}),
		atSubscribe: false,
	}
	s.newReducer(r)
	for _, r := range reducers {
		s.newReducer(r)
	}
	return s
}

func (s *store) newReducer(r reducer) {
	// initial state will be return when current state is nil, so we send nil at here.
	s.GetState[getReducerName(r)] = r(nil, Action{})
	s.reducers = append(s.reducers, r)
}

// Dispatch send action to every reducer
func (s *store) Dispatch(act *Action) {
	s.mu.Lock()
	if s.atSubscribe {
		panic(`you're trying to invoke Dispatch inside the subscribed function`)
	}
	// we dispatch action to every reducer, and reducer update mapping state.
	for _, r := range s.reducers {
		funcName := getReducerName(r)
		s.GetState[funcName] = r(s.GetState[funcName], *act)
	}
	// we call subscribed function after state updated.
	s.atSubscribe = true
	for _, subscribtor := range s.subscribes {
		subscribtor()
	}
	s.atSubscribe = false
	s.mu.Unlock()
}

// Subscribe emit argument into subscribes chain, it will be invoked when Dispatch. !Warning, subscribed function can't invoke Dispatch, it will panic
func (s *store) Subscribe(subscribetor func()) {
	s.subscribes = append(s.subscribes, subscribetor)
}
