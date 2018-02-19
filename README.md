# redux
[![version badges](https://img.shields.io/badge/version-1.2.2-blue.svg)](https://github.com/dannypsnl/redux/releases)
[![Build Status](https://travis-ci.org/dannypsnl/redux.svg?branch=master)](https://travis-ci.org/dannypsnl/redux)
[![Go Report Card](https://goreportcard.com/badge/github.com/dannypsnl/redux)](https://goreportcard.com/report/github.com/dannypsnl/redux)
[![Coverage Status](https://coveralls.io/repos/github/dannypsnl/redux/badge.svg?branch=master)](https://coveralls.io/github/dannypsnl/redux?branch=master)
[![GoDoc](https://godoc.org/github.com/dannypsnl/redux?status.svg)](https://godoc.org/github.com/dannypsnl/redux)
[![GitHub license](https://img.shields.io/github/license/dannypsnl/redux.svg)](https://github.com/dannypsnl/redux/blob/master/LICENSE)
<br>
This is a redux implementation in Go.<br>
I hope this project can help you manage the complex update flow in Go.<br>
[Origin version](https://github.com/reactjs/redux)
## Install
```bash
$ go get github.com/dannypsnl/redux
```
## Usage
### pkgs
#### redux/store
- `New` Create New Store by reducers(at least one reducer)
- `Dispatch` recieve then send action to every reducers to update state<br>
And you should not call `Dispatch` in Subscribetor, you will get dead lock.
- `Subscribe` recieve a func without args will be invoked by every next Dispatch<br>
And you should not call `Subscribe` in Subscribetor, you will get a panic warning.
- `Marshal` return state as JSON format string<br>
#### redux/action
By this module, we can have a better
- `New` recieve a string and return a pointer to Action for you<br>
- `Action` is a type contain `Type` & `Args`<br>
`Type` is just a string help reducer juage what should them do.<br>
`Args` is a `map[string]interface{}` contain a lot values, think about we Dispatch login Action<br>
We need user & password to do this State update, so we will put user & password's value in the Action::Args<br>
Again, only reducer should use Args, so cast is safety.
- `Arg` help you append Argument onto Action with fluent API.<br>
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
