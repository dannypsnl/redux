package action

import (
	"testing"
)

func TestActionArgsInit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(`Args did not be init`)
		}
	}()
	act := New("TEST_MESSAGE")
	act.Args["author"] = "danny"
}

func TestActionFluentAPI(t *testing.T) {
	act := New("TEST_MESSAGE").
		Arg("author", "danny").
		Arg("age", 20)
	if act.Args["author"] != "danny" ||
		act.Args["age"] != 20 {
		t.Error(`Arg API can't work`)
	}
}
