package rematch

import (
	"testing"
)

func Benchmark_getReducerName(b *testing.B) {
	c := NewCountingModel()
	for i := 0; i < b.N; i++ {
		getReducerName(c.Increase)
	}
}
