package redux

import "testing"

func TestSerialize(t *testing.T) {
	store := NewStore(counter)
	store.Dispatch(SendAction("INC"))
	store.Dispatch(SendAction("INC"))
	expected := "{\n  \"counter\":2\n}"
	if store.JSON() != expected {
		t.Errorf("serialized result is not expected, expected: %s, actual: %s", expected, store.JSON())
	}
}
