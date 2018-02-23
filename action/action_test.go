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
	tr := &Tester{t}
	act := New("TEST_MESSAGE").
		Arg("author", "danny").
		Arg("age", 20)
	tr.Assert_EQ("danny", act.Args["author"])
	tr.Assert_EQ(20, act.Args["age"])
}
