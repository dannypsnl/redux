# redux
[![Build Status](https://travis-ci.org/dannypsnl/redux.svg?branch=master)](https://travis-ci.org/dannypsnl/redux)
[![Go Report Card](https://goreportcard.com/badge/github.com/dannypsnl/redux)](https://goreportcard.com/report/github.com/dannypsnl/redux)
[![Coverage Status](https://coveralls.io/repos/github/dannypsnl/redux/badge.svg?branch=master)](https://coveralls.io/github/dannypsnl/redux?branch=master)
[![GoDoc](https://godoc.org/github.com/dannypsnl/redux?status.svg)](https://godoc.org/github.com/dannypsnl/redux)<br>
This is a little try for redux in go.
## Install
```bash
$ go get https://github.com/dannypsnl/redux
```
## Usage
```go
import "github.com/dannypsnl/redux"

func counter(state interface{}, action redux.Action) interface{} {
    // Initial State
    if state == nil {
        return 0
    }
    switch action.Type {
    case "INC":
        return state.(int)+1
    case "DEC":
        return state.(int)-1
    default:
        return state
    }
}

func main() {
    store := redux.NewStore(counter) // counter state be initial at here, it's 0
    // Subscribe's function will be invoke when Dispatch
    store.Subscribe(func() {
        fmt.Printf("Now state is %v\n", store.GetState["counter"])
    })
    store.Dispatch(redux.SendAction("INC")) // state increase by action, now is 1
}
```
