package store

import (
	"reflect"
	"runtime"
	"strings"
)

// Inter comment
// getReducerName is a helper func to get function's ref name.
func getReducerName(r reducer) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
	// fullName's format is `package.function_name`
	// we don't want package part.
	// package is full path(GOPATH/src/package_part) to it
	return fullName[strings.LastIndexByte(fullName, '.')+1:]
}
