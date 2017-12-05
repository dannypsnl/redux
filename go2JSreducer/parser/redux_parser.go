// Generated from Redux.g4 by ANTLR 4.7.

package parser // Redux

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 139,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 6, 2, 24, 10,
	2, 13, 2, 14, 2, 25, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 70,
	10, 5, 13, 5, 14, 5, 71, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 79, 10, 6,
	12, 6, 14, 6, 82, 11, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 99, 10, 8, 12, 8, 14, 8,
	102, 11, 8, 3, 8, 3, 8, 5, 8, 106, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 120, 10, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 129, 10, 11, 3, 11,
	3, 11, 3, 11, 7, 11, 134, 10, 11, 12, 11, 14, 11, 137, 11, 11, 3, 11, 2,
	3, 20, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 3, 3, 2, 27, 30, 2, 141,
	2, 23, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 69, 3, 2, 2,
	2, 10, 73, 3, 2, 2, 2, 12, 86, 3, 2, 2, 2, 14, 105, 3, 2, 2, 2, 16, 107,
	3, 2, 2, 2, 18, 119, 3, 2, 2, 2, 20, 128, 3, 2, 2, 2, 22, 24, 5, 4, 3,
	2, 23, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26,
	3, 2, 2, 2, 26, 3, 3, 2, 2, 2, 27, 30, 5, 6, 4, 2, 28, 30, 5, 18, 10, 2,
	29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2, 31, 32, 7, 3,
	2, 2, 32, 33, 7, 33, 2, 2, 33, 34, 7, 4, 2, 2, 34, 35, 7, 33, 2, 2, 35,
	36, 7, 5, 2, 2, 36, 37, 7, 6, 2, 2, 37, 38, 7, 33, 2, 2, 38, 39, 7, 7,
	2, 2, 39, 40, 7, 8, 2, 2, 40, 41, 7, 5, 2, 2, 41, 42, 7, 9, 2, 2, 42, 43,
	7, 10, 2, 2, 43, 44, 7, 33, 2, 2, 44, 45, 7, 11, 2, 2, 45, 46, 7, 12, 2,
	2, 46, 47, 7, 9, 2, 2, 47, 48, 7, 13, 2, 2, 48, 49, 5, 20, 11, 2, 49, 50,
	7, 14, 2, 2, 50, 51, 7, 15, 2, 2, 51, 52, 7, 33, 2, 2, 52, 53, 7, 16, 2,
	2, 53, 54, 7, 17, 2, 2, 54, 55, 7, 9, 2, 2, 55, 56, 5, 8, 5, 2, 56, 57,
	7, 18, 2, 2, 57, 58, 7, 19, 2, 2, 58, 59, 7, 13, 2, 2, 59, 60, 5, 20, 11,
	2, 60, 61, 7, 14, 2, 2, 61, 62, 7, 14, 2, 2, 62, 7, 3, 2, 2, 2, 63, 64,
	7, 20, 2, 2, 64, 65, 5, 20, 11, 2, 65, 66, 7, 19, 2, 2, 66, 67, 7, 13,
	2, 2, 67, 68, 5, 20, 11, 2, 68, 70, 3, 2, 2, 2, 69, 63, 3, 2, 2, 2, 70,
	71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 9, 3, 2, 2,
	2, 73, 74, 7, 33, 2, 2, 74, 75, 7, 4, 2, 2, 75, 80, 5, 14, 8, 2, 76, 77,
	7, 6, 2, 2, 77, 79, 5, 14, 8, 2, 78, 76, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2,
	80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3,
	2, 2, 2, 83, 84, 7, 8, 2, 2, 84, 85, 5, 14, 8, 2, 85, 11, 3, 2, 2, 2, 86,
	87, 7, 33, 2, 2, 87, 88, 5, 14, 8, 2, 88, 13, 3, 2, 2, 2, 89, 106, 7, 33,
	2, 2, 90, 91, 7, 21, 2, 2, 91, 92, 7, 9, 2, 2, 92, 93, 5, 12, 7, 2, 93,
	94, 7, 14, 2, 2, 94, 106, 3, 2, 2, 2, 95, 96, 7, 22, 2, 2, 96, 100, 7,
	9, 2, 2, 97, 99, 5, 10, 6, 2, 98, 97, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2,
	100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 100,
	3, 2, 2, 2, 103, 106, 7, 14, 2, 2, 104, 106, 3, 2, 2, 2, 105, 89, 3, 2,
	2, 2, 105, 90, 3, 2, 2, 2, 105, 95, 3, 2, 2, 2, 105, 104, 3, 2, 2, 2, 106,
	15, 3, 2, 2, 2, 107, 108, 7, 33, 2, 2, 108, 109, 5, 14, 8, 2, 109, 110,
	7, 23, 2, 2, 110, 111, 5, 20, 11, 2, 111, 17, 3, 2, 2, 2, 112, 113, 7,
	24, 2, 2, 113, 120, 5, 16, 9, 2, 114, 115, 7, 25, 2, 2, 115, 120, 5, 16,
	9, 2, 116, 117, 7, 26, 2, 2, 117, 118, 7, 33, 2, 2, 118, 120, 5, 14, 8,
	2, 119, 112, 3, 2, 2, 2, 119, 114, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 120,
	19, 3, 2, 2, 2, 121, 122, 8, 11, 1, 2, 122, 123, 7, 4, 2, 2, 123, 124,
	5, 20, 11, 2, 124, 125, 7, 8, 2, 2, 125, 129, 3, 2, 2, 2, 126, 129, 7,
	34, 2, 2, 127, 129, 7, 32, 2, 2, 128, 121, 3, 2, 2, 2, 128, 126, 3, 2,
	2, 2, 128, 127, 3, 2, 2, 2, 129, 135, 3, 2, 2, 2, 130, 131, 12, 6, 2, 2,
	131, 132, 9, 2, 2, 2, 132, 134, 5, 20, 11, 7, 133, 130, 3, 2, 2, 2, 134,
	137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 21, 3,
	2, 2, 2, 137, 135, 3, 2, 2, 2, 11, 25, 29, 71, 80, 100, 105, 119, 128,
	135,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "'interface{}'", "','", "'Action'", "')'", "'{'",
	"'if'", "'=='", "'nil'", "'return'", "'}'", "'switch'", "'.'", "'Type'",
	"'default'", "':'", "'case'", "'struct'", "'interface'", "'='", "'const'",
	"'var'", "'type'", "'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "WS", "String", "Ident", "Number",
}

var ruleNames = []string{
	"prog", "stat", "reducer", "cases", "method", "data", "typeFlow", "define",
	"defineGlobal", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ReduxParser struct {
	*antlr.BaseParser
}

func NewReduxParser(input antlr.TokenStream) *ReduxParser {
	this := new(ReduxParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Redux.g4"

	return this
}

// ReduxParser tokens.
const (
	ReduxParserEOF    = antlr.TokenEOF
	ReduxParserT__0   = 1
	ReduxParserT__1   = 2
	ReduxParserT__2   = 3
	ReduxParserT__3   = 4
	ReduxParserT__4   = 5
	ReduxParserT__5   = 6
	ReduxParserT__6   = 7
	ReduxParserT__7   = 8
	ReduxParserT__8   = 9
	ReduxParserT__9   = 10
	ReduxParserT__10  = 11
	ReduxParserT__11  = 12
	ReduxParserT__12  = 13
	ReduxParserT__13  = 14
	ReduxParserT__14  = 15
	ReduxParserT__15  = 16
	ReduxParserT__16  = 17
	ReduxParserT__17  = 18
	ReduxParserT__18  = 19
	ReduxParserT__19  = 20
	ReduxParserT__20  = 21
	ReduxParserT__21  = 22
	ReduxParserT__22  = 23
	ReduxParserT__23  = 24
	ReduxParserT__24  = 25
	ReduxParserT__25  = 26
	ReduxParserT__26  = 27
	ReduxParserT__27  = 28
	ReduxParserWS     = 29
	ReduxParserString = 30
	ReduxParserIdent  = 31
	ReduxParserNumber = 32
)

// ReduxParser rules.
const (
	ReduxParserRULE_prog         = 0
	ReduxParserRULE_stat         = 1
	ReduxParserRULE_reducer      = 2
	ReduxParserRULE_cases        = 3
	ReduxParserRULE_method       = 4
	ReduxParserRULE_data         = 5
	ReduxParserRULE_typeFlow     = 6
	ReduxParserRULE_define       = 7
	ReduxParserRULE_defineGlobal = 8
	ReduxParserRULE_expr         = 9
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *ReduxParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReduxParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ReduxParserT__0)|(1<<ReduxParserT__21)|(1<<ReduxParserT__22)|(1<<ReduxParserT__23))) != 0) {
		{
			p.SetState(20)
			p.Stat()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Reducer() IReducerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReducerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReducerContext)
}

func (s *StatContext) DefineGlobal() IDefineGlobalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineGlobalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineGlobalContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *ReduxParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReduxParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Reducer()
		}

	case ReduxParserT__21, ReduxParserT__22, ReduxParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.DefineGlobal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReducerContext is an interface to support dynamic dispatch.
type IReducerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInitialState returns the initialState rule contexts.
	GetInitialState() IExprContext

	// SetInitialState sets the initialState rule contexts.
	SetInitialState(IExprContext)

	// IsReducerContext differentiates from other interfaces.
	IsReducerContext()
}

type ReducerContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	initialState IExprContext
}

func NewEmptyReducerContext() *ReducerContext {
	var p = new(ReducerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_reducer
	return p
}

func (*ReducerContext) IsReducerContext() {}

func NewReducerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReducerContext {
	var p = new(ReducerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_reducer

	return p
}

func (s *ReducerContext) GetParser() antlr.Parser { return s.parser }

func (s *ReducerContext) GetInitialState() IExprContext { return s.initialState }

func (s *ReducerContext) SetInitialState(v IExprContext) { s.initialState = v }

func (s *ReducerContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(ReduxParserIdent)
}

func (s *ReducerContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, i)
}

func (s *ReducerContext) Cases() ICasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *ReducerContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ReducerContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReducerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReducerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReducerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterReducer(s)
	}
}

func (s *ReducerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitReducer(s)
	}
}

func (p *ReduxParser) Reducer() (localctx IReducerContext) {
	localctx = NewReducerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ReduxParserRULE_reducer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(ReduxParserT__0)
	}
	{
		p.SetState(30)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(31)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(32)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(33)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(34)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(35)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(36)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(37)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(38)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(39)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(40)
		p.Match(ReduxParserT__7)
	}
	{
		p.SetState(41)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(42)
		p.Match(ReduxParserT__8)
	}
	{
		p.SetState(43)
		p.Match(ReduxParserT__9)
	}
	{
		p.SetState(44)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(45)
		p.Match(ReduxParserT__10)
	}
	{
		p.SetState(46)

		var _x = p.expr(0)

		localctx.(*ReducerContext).initialState = _x
	}
	{
		p.SetState(47)
		p.Match(ReduxParserT__11)
	}
	{
		p.SetState(48)
		p.Match(ReduxParserT__12)
	}
	{
		p.SetState(49)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(50)
		p.Match(ReduxParserT__13)
	}
	{
		p.SetState(51)
		p.Match(ReduxParserT__14)
	}
	{
		p.SetState(52)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(53)
		p.Cases()
	}
	{
		p.SetState(54)
		p.Match(ReduxParserT__15)
	}
	{
		p.SetState(55)
		p.Match(ReduxParserT__16)
	}
	{
		p.SetState(56)
		p.Match(ReduxParserT__10)
	}
	{
		p.SetState(57)
		p.expr(0)
	}
	{
		p.SetState(58)
		p.Match(ReduxParserT__11)
	}
	{
		p.SetState(59)
		p.Match(ReduxParserT__11)
	}

	return localctx
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_cases
	return p
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CasesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitCases(s)
	}
}

func (p *ReduxParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ReduxParserRULE_cases)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ReduxParserT__17 {
		{
			p.SetState(61)
			p.Match(ReduxParserT__17)
		}
		{
			p.SetState(62)
			p.expr(0)
		}
		{
			p.SetState(63)
			p.Match(ReduxParserT__16)
		}
		{
			p.SetState(64)
			p.Match(ReduxParserT__10)
		}
		{
			p.SetState(65)
			p.expr(0)
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) Ident() antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, 0)
}

func (s *MethodContext) AllTypeFlow() []ITypeFlowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem())
	var tst = make([]ITypeFlowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFlowContext)
		}
	}

	return tst
}

func (s *MethodContext) TypeFlow(i int) ITypeFlowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFlowContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *ReduxParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ReduxParserRULE_method)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(72)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(73)
		p.TypeFlow()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ReduxParserT__3 {
		{
			p.SetState(74)
			p.Match(ReduxParserT__3)
		}
		{
			p.SetState(75)
			p.TypeFlow()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(82)
		p.TypeFlow()
	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) Ident() antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, 0)
}

func (s *DataContext) TypeFlow() ITypeFlowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFlowContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *ReduxParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ReduxParserRULE_data)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(85)
		p.TypeFlow()
	}

	return localctx
}

// ITypeFlowContext is an interface to support dynamic dispatch.
type ITypeFlowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFlowContext differentiates from other interfaces.
	IsTypeFlowContext()
}

type TypeFlowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFlowContext() *TypeFlowContext {
	var p = new(TypeFlowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_typeFlow
	return p
}

func (*TypeFlowContext) IsTypeFlowContext() {}

func NewTypeFlowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFlowContext {
	var p = new(TypeFlowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_typeFlow

	return p
}

func (s *TypeFlowContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFlowContext) Ident() antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, 0)
}

func (s *TypeFlowContext) Data() IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *TypeFlowContext) AllMethod() []IMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodContext)(nil)).Elem())
	var tst = make([]IMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodContext)
		}
	}

	return tst
}

func (s *TypeFlowContext) Method(i int) IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *TypeFlowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFlowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFlowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterTypeFlow(s)
	}
}

func (s *TypeFlowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitTypeFlow(s)
	}
}

func (p *ReduxParser) TypeFlow() (localctx ITypeFlowContext) {
	localctx = NewTypeFlowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ReduxParserRULE_typeFlow)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(ReduxParserIdent)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(ReduxParserT__18)
		}
		{
			p.SetState(89)
			p.Match(ReduxParserT__6)
		}
		{
			p.SetState(90)
			p.Data()
		}
		{
			p.SetState(91)
			p.Match(ReduxParserT__11)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(ReduxParserT__19)
		}
		{
			p.SetState(94)
			p.Match(ReduxParserT__6)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ReduxParserIdent {
			{
				p.SetState(95)
				p.Method()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(ReduxParserT__11)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)

	}

	return localctx
}

// IDefineContext is an interface to support dynamic dispatch.
type IDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineContext differentiates from other interfaces.
	IsDefineContext()
}

type DefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineContext() *DefineContext {
	var p = new(DefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_define
	return p
}

func (*DefineContext) IsDefineContext() {}

func NewDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineContext {
	var p = new(DefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_define

	return p
}

func (s *DefineContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineContext) Ident() antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, 0)
}

func (s *DefineContext) TypeFlow() ITypeFlowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFlowContext)
}

func (s *DefineContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterDefine(s)
	}
}

func (s *DefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitDefine(s)
	}
}

func (p *ReduxParser) Define() (localctx IDefineContext) {
	localctx = NewDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ReduxParserRULE_define)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(106)
		p.TypeFlow()
	}
	{
		p.SetState(107)
		p.Match(ReduxParserT__20)
	}
	{
		p.SetState(108)
		p.expr(0)
	}

	return localctx
}

// IDefineGlobalContext is an interface to support dynamic dispatch.
type IDefineGlobalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineGlobalContext differentiates from other interfaces.
	IsDefineGlobalContext()
}

type DefineGlobalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineGlobalContext() *DefineGlobalContext {
	var p = new(DefineGlobalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_defineGlobal
	return p
}

func (*DefineGlobalContext) IsDefineGlobalContext() {}

func NewDefineGlobalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineGlobalContext {
	var p = new(DefineGlobalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_defineGlobal

	return p
}

func (s *DefineGlobalContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineGlobalContext) Define() IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *DefineGlobalContext) Ident() antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, 0)
}

func (s *DefineGlobalContext) TypeFlow() ITypeFlowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFlowContext)
}

func (s *DefineGlobalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineGlobalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineGlobalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterDefineGlobal(s)
	}
}

func (s *DefineGlobalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitDefineGlobal(s)
	}
}

func (p *ReduxParser) DefineGlobal() (localctx IDefineGlobalContext) {
	localctx = NewDefineGlobalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ReduxParserRULE_defineGlobal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(ReduxParserT__21)
		}
		{
			p.SetState(111)
			p.Define()
		}

	case ReduxParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(ReduxParserT__22)
		}
		{
			p.SetState(113)
			p.Define()
		}

	case ReduxParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)
			p.Match(ReduxParserT__23)
		}
		{
			p.SetState(115)
			p.Match(ReduxParserIdent)
		}
		{
			p.SetState(116)
			p.TypeFlow()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Number() antlr.TerminalNode {
	return s.GetToken(ReduxParserNumber, 0)
}

func (s *ExprContext) String() antlr.TerminalNode {
	return s.GetToken(ReduxParserString, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *ReduxParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ReduxParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ReduxParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__1:
		{
			p.SetState(120)
			p.Match(ReduxParserT__1)
		}
		{
			p.SetState(121)
			p.expr(0)
		}
		{
			p.SetState(122)
			p.Match(ReduxParserT__5)
		}

	case ReduxParserNumber:
		{
			p.SetState(124)
			p.Match(ReduxParserNumber)
		}

	case ReduxParserString:
		{
			p.SetState(125)
			p.Match(ReduxParserString)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ReduxParserRULE_expr)
			p.SetState(128)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			p.SetState(129)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ReduxParserT__24)|(1<<ReduxParserT__25)|(1<<ReduxParserT__26)|(1<<ReduxParserT__27))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(130)
				p.expr(5)
			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ReduxParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ReduxParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
