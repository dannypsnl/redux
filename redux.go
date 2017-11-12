package redux

type Action struct {
	Type string
}

func SendAction(typ string) *Action {
	return &Action{
		Type: typ,
	}
}

type Reducer func(interface{}, Action) interface{}

type store struct {
	reducer Reducer
	state   interface{}
}

func NewStore(r Reducer) *store {
	return &store{
		reducer: r,
		state:   0,
	}
}

func (s *store) Dispatch(act *Action) {
	s.state = s.reducer(s.state, *act)
}
func (s *store) GetState() interface{} {
	return s.state
}
