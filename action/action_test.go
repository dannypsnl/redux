package action

import (
	"testing"
)

func TestArgsHadBeenInit(t *testing.T) {
	defer func() {
		// If panic, we can recover it
		// So we can know Args didn't be initialize
		if r := recover(); r != nil {
			t.Error(`Args did not be init`)
		}
	}()
	act := New("TEST_MESSAGE")
	// If didn't, this operation will cause panic
	act.Args["author"] = "danny"
}

func TestArgCanWork(t *testing.T) {
	act := New("TEST_MESSAGE").
		Arg("author", "danny").
		Arg("age", 20)
	if act.Args["author"] != "danny" ||
		act.Args["age"] != 20 {
		t.Error(`Arg API can't work`)
	}
}
