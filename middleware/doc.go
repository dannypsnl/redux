// Package middleware provides middleware helper type for redux
//
// Example:
//  func Logger(s *store.Store) middleware.Middleware {
//      return func(next middleware.Next) middleware.Next {
//          return func(act *action.Action) *action.Action {
//              return next(act)
//          }
//      }
//  }
package middleware
