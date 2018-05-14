# redux
[![version badges](https://img.shields.io/badge/version-1.5.1-blue.svg)](https://github.com/dannypsnl/redux/releases)
[![Build Status](https://travis-ci.org/dannypsnl/redux.svg?branch=master)](https://travis-ci.org/dannypsnl/redux)
[![Build status](https://ci.appveyor.com/api/projects/status/l2cqebl1svcgyrpo?svg=true)](https://ci.appveyor.com/project/dannypsnl/redux)
[![Go Report Card](https://goreportcard.com/badge/github.com/dannypsnl/redux)](https://goreportcard.com/report/github.com/dannypsnl/redux)
[![Coverage Status](https://coveralls.io/repos/github/dannypsnl/redux/badge.svg?branch=master)](https://coveralls.io/github/dannypsnl/redux?branch=master)
[![GoDoc](https://godoc.org/github.com/dannypsnl/redux?status.svg)](https://godoc.org/github.com/dannypsnl/redux)
[![GitHub license](https://img.shields.io/github/license/dannypsnl/redux.svg)](https://github.com/dannypsnl/redux/blob/master/LICENSE)

This is a redux implementation in Go.<br>
I hope this project can help you manage the complex update flow in Go.<br>
[Origin version](https://github.com/reactjs/redux)
## Install
```bash
$ go get github.com/dannypsnl/redux
```
## Usage

pkgs       | reference
-----------|--------------------------------------------------------------------------------------------------
store      | [![store](https://godoc.org/github.com/dannypsnl/redux/store?status.svg)](https://godoc.org/github.com/dannypsnl/redux/store)
action     | [![action](https://godoc.org/github.com/dannypsnl/redux/action?status.svg)](https://godoc.org/github.com/dannypsnl/redux/action)
middleware | [![middleware](https://godoc.org/github.com/dannypsnl/redux/middleware?status.svg)](https://godoc.org/github.com/dannypsnl/redux/middleware)
rematch    | [![rematch](https://godoc.org/github.com/dannypsnl/redux/rematch?status.svg)](https://godoc.org/github.com/dannypsnl/redux/rematch)

### Example
[Examples](https://github.com/dannypsnl/redux/tree/master/example)
##### Basic Example
> Ignore other packages
```go
import(
    "github.com/dannypsnl/redux/store"
    "github.com/dannypsnl/redux/action"
)

func counter(state interface{}, action action.Action) interface{} {
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
    store := store.New(counter) // counter state be initial at here, it's 0
    // Subscribe's function will be invoke when Dispatch
    store.Subscribe(func() {
        fmt.Printf("Now state is %v\n", store.GetState("counter"))
        fmt.Printf("%s\n", store.Marshal()) // `{ counter: 1 }`, 1 should be current state, Let's print out the json format of our store
    })
    store.Dispatch(action.New("INC")) // state increase by action, now is 1
}
```
## Changelogs
[Changelogs](https://github.com/dannypsnl/redux/blob/master/CHANGELOG.md)
