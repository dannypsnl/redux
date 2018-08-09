package rematch

import (
	"testing"
)

func BenchmarkGetReducerName(b *testing.B) {
	c := NewCountingModel()
	for i := 0; i < b.N; i++ {
		getReducerName(c.Increase)
	}
}
