# redux

[![version badges](https://img.shields.io/badge/version-2.0.0-blue.svg)](https://github.com/dannypsnl/redux/releases)
[![Build Status](https://travis-ci.org/dannypsnl/redux.svg?branch=master)](https://travis-ci.org/dannypsnl/redux)
[![Build status](https://ci.appveyor.com/api/projects/status/l2cqebl1svcgyrpo?svg=true)](https://ci.appveyor.com/project/dannypsnl/redux)
[![Go Report Card](https://goreportcard.com/badge/github.com/dannypsnl/redux)](https://goreportcard.com/report/github.com/dannypsnl/redux)
[![Coverage Status](https://coveralls.io/repos/github/dannypsnl/redux/badge.svg?branch=master)](https://coveralls.io/github/dannypsnl/redux?branch=master)
[![GoDoc](https://godoc.org/github.com/dannypsnl/redux?status.svg)](https://godoc.org/github.com/dannypsnl/redux)
[![GitHub license](https://img.shields.io/github/license/dannypsnl/redux.svg)](https://github.com/dannypsnl/redux/blob/master/LICENSE)

This is a redux implementation in Go.

I hope this project can help you manage the complex update flow in Go.

[Origin version](https://github.com/reactjs/redux)

## Install

```bash
$ go get github.com/dannypsnl/redux
```

## Usage

Example:

```go
// As you can see, you don't need to checking nil or not now
// Now redux-go will initial the state with zero value of that type
func counter(state int, action string) int {
    switch action {
    case "INC":
        return state + 1
    case "DEC":
        return state - 1
    default:
        return state
    }
}

func main() {
    store := store.New(counter)
    store.Dispatch("INC")
    fmt.Printf("%d\n", store.GetState("counter")) // should print out: 1
}
```
