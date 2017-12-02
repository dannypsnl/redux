package redux

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

type store struct {
	// reducers contains every reducer of this store.
	reducers []Reducer
	// GetState contain key map state
	// and we use key to spread reducer's state
	// for example: "counter" mapping to reducer named counter.
	// when use GetState["counter"], we will got the current state of "counter" key's mapping target.
	GetState   map[string]interface{}
	subscribes []func()
}

func NewStore(reducer Reducer, reducers ...Reducer) *store {
	s := &store{
		GetState: make(map[string]interface{}),
	}
	s.newReducer(reducer)
	for _, reducer := range reducers {
		s.newReducer(reducer)
	}
	return s
}

func (s *store) newReducer(reducer Reducer) {
	s.GetState[getReducerName(reducer)] = reducer(nil, Action{})
	s.reducers = append(s.reducers, reducer)
}

func (s *store) Dispatch(act *Action) {
	// we dispatch action to every reducer, and reducer update mapping state.
	for _, reducer := range s.reducers {
		func_name := getReducerName(reducer)
		s.GetState[func_name] = reducer(s.GetState[func_name], *act)
	}
	// we call subscribed function after state updated.
	for _, subscribtor := range s.subscribes {
		subscribtor()
	}
}

func (s *store) Subscribe(subscribe_fn func()) {
	s.subscribes = append(s.subscribes, subscribe_fn)
}
