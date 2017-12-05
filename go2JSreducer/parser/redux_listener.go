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

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterReducer is called when entering the reducer production.
	EnterReducer(c *ReducerContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

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

	// EnterPackageStat is called when entering the packageStat production.
	EnterPackageStat(c *PackageStatContext)

	// EnterImportStat is called when entering the importStat production.
	EnterImportStat(c *ImportStatContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitReducer is called when exiting the reducer production.
	ExitReducer(c *ReducerContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

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

	// ExitPackageStat is called when exiting the packageStat production.
	ExitPackageStat(c *PackageStatContext)

	// ExitImportStat is called when exiting the importStat production.
	ExitImportStat(c *ImportStatContext)
}
