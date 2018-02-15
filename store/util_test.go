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
			t.Errorf(`getReducerName didn't get correct name, expected: %s, actual %s`, expected, getReducerName(reducer))
		}
	}
}
