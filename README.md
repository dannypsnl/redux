# redux

[![version badges](https://img.shields.io/badge/version-2.2.2-blue.svg)](https://github.com/dannypsnl/redux/releases)
![Build Status](https://github.com/dannypsnl/redux/workflows/Go/badge.svg?branch=master)
[![Build status](https://ci.appveyor.com/api/projects/status/l2cqebl1svcgyrpo?svg=true)](https://ci.appveyor.com/project/dannypsnl/redux)
[![codecov](https://codecov.io/gh/dannypsnl/redux/branch/master/graph/badge.svg)](https://codecov.io/gh/dannypsnl/redux)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dannypsnl/redux/v2)
[![GitHub license](https://img.shields.io/github/license/dannypsnl/redux.svg)](https://github.com/dannypsnl/redux/blob/master/LICENSE)

This is a redux implementation in Go.

[Origin version](https://github.com/reactjs/redux)

## Install

```bash
$ go get github.com/dannypsnl/redux
```

## Usage

```go
import (
    // must import, else go module would think following package is a module
    _ "github.com/dannypsnl/redux/v2"
    // v2 store
    "github.com/dannypsnl/redux/v2/store"
    // v2 rematch(optional)
    "github.com/dannypsnl/redux/v2/rematch"
)
```

Basic example

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
    // redux/v2/store
    store := store.New(counter)
    store.Dispatch("INC")
    fmt.Printf("%d\n", store.StateOf(counter)) // should print out: 1
}
```

More stunning(lambda can work)

```go
func main() {
    counter := func(state, payload int) int {
        return state + payload
    }
    store := store.New(counter)
    store.Dispatch(100)
    store.Dispatch(-50)
    fmt.Printf("%d\n", store.StateOf(counter)) // should print out: 50
}
```

### Rematch

And more...

```go
type CountingModel struct {
    rematch.Reducer
    State int
}

func (cm *CountingModel) Increase(s, payload int) int {
    return s + payload
}

func (cm *CountingModel) Decrease(s, payload int) int {
    return s - payload
}

func main() {
    c := &CountingModel{
        State: 0,
    }
    store := store.New(c)
    store.Dispatch(c.Action(c.Increase).With(100))
    store.Dispatch(c.Action(c.Decrease).With(50))

    fmt.Printf("result: %d\n", store.StateOf(c)) // expect: 50
}
```

Then let's have a party

```go
type CountingModel struct {
    Reducer
    State int

    Increase *rematch.Action `action:"IncreaseImpl"`
}

func (c *CountingModel) IncreaseImpl(s, payload int) int {
    return s + payload
}

func main() {
    c := &CountingModel {
        State: 0,
    }
    store := store.New(c)
    store.Dispatch(c.Increase.With(30))
    store.Dispatch(c.Increase.With(20))

    fmt.Printf("result: %d\n", store.StateOf(c)) // expect: 50
}
```
