package store

import (
	"reflect"
	"runtime"
	"strings"
)

// Inter comment
// getReducerName is a helper func to get function's ref name.
func getReducerName(r reducer) string {
	// this code will get package.function_name, so we have to drop package part.
	fullName := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
	return fullName[strings.LastIndexByte(fullName, '.')+1:]
}
