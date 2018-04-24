// rematch provide rematch API
//
// Example:
//   counter := rematch.Reducer {
//       State: 0,
//       Reducers: rematch.Reducers {
//           "INC": func(state interface{}, act action.Action) interface{} {
//                return state.(int)+act.Args["payload"].(int)
//            },
//            "DEC": func(state interface{}, act action.Action) interface{} {
//                return state.(int)-act.Args["payload"].(int)
//            },
//       }
//   }
//   store := store.New(counter)
//   store.Dispatch(counter.
//       Action("INC").
//           Arg("payload", 10))
package rematch
