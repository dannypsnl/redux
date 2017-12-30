package redux

import (
	"testing"
)

func TestActionArgsInit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(`Args did not be init`)
		}
	}()
	act := SendAction("TEST_MESSAGE")
	act.Args["author"] = "danny"
}
