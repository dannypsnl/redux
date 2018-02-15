package action

import "testing"

type Tester struct {
	t *testing.T
}

func (t *Tester) Assert_EQ(expected, actual interface{}) {
	if expected != actual {
		t.t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
