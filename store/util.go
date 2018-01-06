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
	// this code will get package.function_name
	// so we have to drop package part.
	// package is full path to it
	return fullName[strings.LastIndexByte(fullName, '.')+1:]
}
