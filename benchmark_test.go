package redux

import (
	"testing"
	"time"
)

func BenchmarkDispatch(b *testing.B) {
	store := NewStore(counter, jump)
	store.Subscribe(func() {
		time.Sleep(10 * time.Millisecond)
	})
	store.Subscribe(func() {
		time.Sleep(10 * time.Millisecond)
	})
	for i := 0; i < b.N; i++ {
		store.Dispatch(SendAction("JUMP"))
	}
}
