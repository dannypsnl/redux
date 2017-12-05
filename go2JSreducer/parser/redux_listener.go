// Generated from Redux.g4 by ANTLR 4.7.

package parser // Redux

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReduxListener is a complete listener for a parse tree produced by ReduxParser.
type ReduxListener interface {
	antlr.ParseTreeListener

	// EnterReducer is called when entering the reducer production.
	EnterReducer(c *ReducerContext)

	// ExitReducer is called when exiting the reducer production.
	ExitReducer(c *ReducerContext)
}
