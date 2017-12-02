package redux

import (
	"reflect"
	"runtime"
	"strings"
)

// getReducerName is a helper func to get function's ref name.
func getReducerName(r Reducer) string {
	// this code will get package.function_name, so we have to drop package part.
	full_name := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
	return full_name[strings.LastIndexByte(full_name, '.')+1:]
}
