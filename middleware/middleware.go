// Package middleware provides middleware helper type for redux
package middleware

import (
	"github.com/dannypsnl/redux/action"
)

// Middleware is exist to let user create middleware code much more faster
type Middleware func(Next) Next

// Next is exist to let user create middleware code much more faster
type Next func(*action.Action) *action.Action
