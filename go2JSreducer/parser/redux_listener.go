// Generated from Redux.g4 by ANTLR 4.7.

package parser // Redux

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReduxListener is a complete listener for a parse tree produced by ReduxParser.
type ReduxListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterReducer is called when entering the reducer production.
	EnterReducer(c *ReducerContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitReducer is called when exiting the reducer production.
	ExitReducer(c *ReducerContext)
}
