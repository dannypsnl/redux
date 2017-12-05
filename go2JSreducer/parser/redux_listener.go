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

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterTypeFlow is called when entering the typeFlow production.
	EnterTypeFlow(c *TypeFlowContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterDefineGlobal is called when entering the defineGlobal production.
	EnterDefineGlobal(c *DefineGlobalContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitReducer is called when exiting the reducer production.
	ExitReducer(c *ReducerContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitTypeFlow is called when exiting the typeFlow production.
	ExitTypeFlow(c *TypeFlowContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitDefineGlobal is called when exiting the defineGlobal production.
	ExitDefineGlobal(c *DefineGlobalContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
