# redux
This is a little try for redux in go.
## Install
```bash
$ go get https://github.com/dannypsnl/redux
```
## Usage
```go
import "github.com/dannypsnl/redux"

func counter(state interface{}, act redux.Action) interface{} {
    switch act.Type {
    case "INC":
        return state.(int)+1
    case "DEC":
        return state.(int)-1
    default:
        return state
    }
}

func main() {
    store := redux.NewStore(counter)
    store.Dispatch(redux.SendAction("INC"))

    fmt.Printf("Now state is %v\n", store.GetState())
}
```
