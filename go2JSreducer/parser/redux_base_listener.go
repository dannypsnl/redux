// Generated from Redux.g4 by ANTLR 4.7.

package parser // Redux

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReduxListener is a complete listener for a parse tree produced by ReduxParser.
type BaseReduxListener struct{}

var _ ReduxListener = &BaseReduxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReduxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReduxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReduxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReduxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseReduxListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseReduxListener) ExitProg(ctx *ProgContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseReduxListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseReduxListener) ExitStat(ctx *StatContext) {}

// EnterReducer is called when production reducer is entered.
func (s *BaseReduxListener) EnterReducer(ctx *ReducerContext) {}

// ExitReducer is called when production reducer is exited.
func (s *BaseReduxListener) ExitReducer(ctx *ReducerContext) {}

// EnterCases is called when production cases is entered.
func (s *BaseReduxListener) EnterCases(ctx *CasesContext) {}

// ExitCases is called when production cases is exited.
func (s *BaseReduxListener) ExitCases(ctx *CasesContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseReduxListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseReduxListener) ExitMethod(ctx *MethodContext) {}

// EnterData is called when production data is entered.
func (s *BaseReduxListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseReduxListener) ExitData(ctx *DataContext) {}

// EnterTypeFlow is called when production typeFlow is entered.
func (s *BaseReduxListener) EnterTypeFlow(ctx *TypeFlowContext) {}

// ExitTypeFlow is called when production typeFlow is exited.
func (s *BaseReduxListener) ExitTypeFlow(ctx *TypeFlowContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseReduxListener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseReduxListener) ExitDefine(ctx *DefineContext) {}

// EnterDefineGlobal is called when production defineGlobal is entered.
func (s *BaseReduxListener) EnterDefineGlobal(ctx *DefineGlobalContext) {}

// ExitDefineGlobal is called when production defineGlobal is exited.
func (s *BaseReduxListener) ExitDefineGlobal(ctx *DefineGlobalContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseReduxListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseReduxListener) ExitExpr(ctx *ExprContext) {}
