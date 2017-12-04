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

// Store is a type manage our data
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
	// atSubscribe make sure we can't call Dispatch in Subscribed function
	atSubscribe bool
	// We use mu to Lock each Dispatch call
	mu sync.Mutex
}

// NewStore create a Store by reducers
func NewStore(r reducer, reducers ...reducer) *Store {
	s := &Store{
		state:       make(map[string]interface{}),
		atSubscribe: false,
	}
	s.newReducer(r)
	for _, r := range reducers {
		s.newReducer(r)
	}
	return s
}

func (s *Store) newReducer(r reducer) {
	// initial state will be return when current state is nil, so we send nil at here.
	s.state[getReducerName(r)] = r(nil, Action{})
	s.reducers = append(s.reducers, r)
}

// GetState recieve string key to get the mapped state
func (s *Store) GetState(key string) interface{} {
	return s.state[key]
}

// Dispatch send action to every reducer
func (s *Store) Dispatch(act *Action) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.atSubscribe {
		panic(`you're trying to invoke Dispatch inside the subscribed function`)
	}
	// we dispatch action to every reducer, and reducer update mapping state.
	for _, r := range s.reducers {
		funcName := getReducerName(r)
		s.state[funcName] = r(s.state[funcName], *act)
	}
	// we call subscribed function after state updated.
	s.atSubscribe = true
	for _, subscribtor := range s.subscribes {
		subscribtor()
	}
	s.atSubscribe = false
}

// Subscribe emit argument into subscribes chain, it will be invoked when Dispatch. !Warning, subscribed function can't invoke Dispatch, it will panic
func (s *Store) Subscribe(subscribetor func()) {
	s.subscribes = append(s.subscribes, subscribetor)
}
