package redux

import (
	"testing"
)

func TestGetReducerName(t *testing.T) {
	var testReducersName = map[string]reducer{
		`counter`: counter,
		`jump`:    jump,
	}
	for k, v := range testReducersName {
		if getReducerName(v) != k {
			t.Error(`getReducerName didn't get correct name`)
		}
	}
}
