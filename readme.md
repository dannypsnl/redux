# redux
[![Build Status](https://travis-ci.org/dannypsnl/redux.svg?branch=master)](https://travis-ci.org/dannypsnl/redux)
[![Go Report Card](https://goreportcard.com/badge/github.com/dannypsnl/redux)](https://goreportcard.com/report/github.com/dannypsnl/redux)
[![Coverage Status](https://coveralls.io/repos/github/dannypsnl/redux/badge.svg)](https://coveralls.io/github/dannypsnl/redux)
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
        fmt.Printf("Now state is %v\n", store.GetState("counter"))
    })
    store.Dispatch(redux.SendAction("INC")) // state increase by action, now is 1
}
```
- NewStore
Create New Store by reducers(at least one reducer)
- Dispatch
Dispatch recieve then send your action object to every reducers to update state
- SendAction
SendAction recieve a string and return a pointer to Action for you<br>
The reason for it is because we usually only need Type, so SendAction reduce the code for you and reduce the opportunity make fault<br>
- Action
Action is a type contain Type & Args<br>
Type is just a string help reducer juage what should them do.<br>
Args is a map[string]interface{} contain a lot values, think about we Dispatch login Action<br>
We need user & password to do this State update, so we will put user & password's value in the Action::Args<br>
Again, only reducer should use Args, so cast is safety
- Subscribe
Subscribe recieve a func without args will be invoked by every next Dispatch
- JSON
JSON return state as JSON format string