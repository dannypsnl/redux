package store

import (
	"testing"
)

func TestGetReducerName(t *testing.T) {
	var testReducersName = map[string]reducer{
		`counter`: counter,
		`jump`:    jump,
	}
	for expected, reducer := range testReducersName {
		if getReducerName(reducer) != expected {
			t.Error(`getReducerName didn't get correct name`)
		}
	}
}
