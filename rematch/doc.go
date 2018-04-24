// rematch provide rematch API
//
// Example:
//   func counter(s interface{}, a action.Action) interface{} {
//       return reCounter.Reduce()(s, a)
//   }
//   func init() {
//       reCounter = rematch.Reducer {
//           State: 0,
//           Reducers: rematch.Reducers {
//               "INC": func(state interface{}, act action.Action) interface{} {
//                   return state.(int)+act.Args["payload"].(int)
//               },
//               "DEC": func(state interface{}, act action.Action) interface{} {
//                   return state.(int)-act.Args["payload"].(int)
//               },
//           }
//       }
//   }
//   func main() {
//       store := store.New(counter)
//       store.Dispatch(reCounter.
//           Action("INC").
//               Arg("payload", 10))
//   }
package rematch
