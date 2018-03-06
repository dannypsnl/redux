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

import (
	"github.com/dannypsnl/redux/action"
)

// Middleware is exist to let user create middleware code much more faster
type Middleware func(Next) Next

// Next is exist to let user create middleware code much more faster
type Next func(*action.Action) *action.Action
