package rematch

import (
	"testing"
)

type CountingModel struct {
	Reducer
	State int
}

func (cm *CountingModel) Increase(state int, payload int) int {
	return state + payload
}

func BenchmarkGetReducerName(b *testing.B) {
	c := &CountingModel{}
	for i := 0; i < b.N; i++ {
		getReducerName(c.Increase)
	}
}
