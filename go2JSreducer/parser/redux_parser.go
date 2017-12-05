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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 108,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14,
	2, 23, 3, 3, 3, 3, 5, 3, 28, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 7, 5, 49, 10, 5, 12, 5, 14, 5, 52, 11, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7,
	69, 10, 7, 12, 7, 14, 7, 72, 11, 7, 3, 7, 3, 7, 5, 7, 76, 10, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	90, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 98, 10, 10,
	3, 10, 3, 10, 3, 10, 7, 10, 103, 10, 10, 12, 10, 14, 10, 106, 11, 10, 3,
	10, 2, 3, 18, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3, 3, 2, 17, 20, 2,
	109, 2, 21, 3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8, 43, 3,
	2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 75, 3, 2, 2, 2, 14, 77, 3, 2, 2, 2, 16,
	89, 3, 2, 2, 2, 18, 97, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2,
	2, 22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3,
	2, 2, 2, 25, 28, 5, 6, 4, 2, 26, 28, 5, 16, 9, 2, 27, 25, 3, 2, 2, 2, 27,
	26, 3, 2, 2, 2, 28, 5, 3, 2, 2, 2, 29, 30, 7, 3, 2, 2, 30, 31, 7, 22, 2,
	2, 31, 32, 7, 4, 2, 2, 32, 33, 7, 22, 2, 2, 33, 34, 7, 5, 2, 2, 34, 35,
	7, 6, 2, 2, 35, 36, 7, 22, 2, 2, 36, 37, 7, 7, 2, 2, 37, 38, 7, 8, 2, 2,
	38, 39, 7, 5, 2, 2, 39, 40, 7, 9, 2, 2, 40, 41, 7, 22, 2, 2, 41, 42, 7,
	10, 2, 2, 42, 7, 3, 2, 2, 2, 43, 44, 7, 22, 2, 2, 44, 45, 7, 4, 2, 2, 45,
	50, 5, 12, 7, 2, 46, 47, 7, 6, 2, 2, 47, 49, 5, 12, 7, 2, 48, 46, 3, 2,
	2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 53,
	3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 8, 2, 2, 54, 55, 5, 12, 7, 2,
	55, 9, 3, 2, 2, 2, 56, 57, 7, 22, 2, 2, 57, 58, 5, 12, 7, 2, 58, 11, 3,
	2, 2, 2, 59, 76, 7, 22, 2, 2, 60, 61, 7, 11, 2, 2, 61, 62, 7, 9, 2, 2,
	62, 63, 5, 10, 6, 2, 63, 64, 7, 10, 2, 2, 64, 76, 3, 2, 2, 2, 65, 66, 7,
	12, 2, 2, 66, 70, 7, 9, 2, 2, 67, 69, 5, 8, 5, 2, 68, 67, 3, 2, 2, 2, 69,
	72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2,
	2, 72, 70, 3, 2, 2, 2, 73, 76, 7, 10, 2, 2, 74, 76, 3, 2, 2, 2, 75, 59,
	3, 2, 2, 2, 75, 60, 3, 2, 2, 2, 75, 65, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2,
	76, 13, 3, 2, 2, 2, 77, 78, 7, 22, 2, 2, 78, 79, 5, 12, 7, 2, 79, 80, 7,
	13, 2, 2, 80, 81, 5, 18, 10, 2, 81, 15, 3, 2, 2, 2, 82, 83, 7, 14, 2, 2,
	83, 90, 5, 14, 8, 2, 84, 85, 7, 15, 2, 2, 85, 90, 5, 14, 8, 2, 86, 87,
	7, 16, 2, 2, 87, 88, 7, 22, 2, 2, 88, 90, 5, 12, 7, 2, 89, 82, 3, 2, 2,
	2, 89, 84, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2, 90, 17, 3, 2, 2, 2, 91, 92,
	8, 10, 1, 2, 92, 93, 7, 4, 2, 2, 93, 94, 5, 18, 10, 2, 94, 95, 7, 8, 2,
	2, 95, 98, 3, 2, 2, 2, 96, 98, 7, 23, 2, 2, 97, 91, 3, 2, 2, 2, 97, 96,
	3, 2, 2, 2, 98, 104, 3, 2, 2, 2, 99, 100, 12, 5, 2, 2, 100, 101, 9, 2,
	2, 2, 101, 103, 5, 18, 10, 6, 102, 99, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2,
	104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 19, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 10, 23, 27, 50, 70, 75, 89, 97, 104,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "'interface{}'", "','", "'Action'", "')'", "'{'",
	"'}'", "'struct'", "'interface'", "'='", "'const'", "'var'", "'type'",
	"'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "WS", "Ident", "Number",
}

var ruleNames = []string{
	"prog", "stat", "reducer", "method", "data", "typeFlow", "define", "defineGlobal",
	"expr",
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
	ReduxParserWS     = 19
	ReduxParserIdent  = 20
	ReduxParserNumber = 21
)

// ReduxParser rules.
const (
	ReduxParserRULE_prog         = 0
	ReduxParserRULE_stat         = 1
	ReduxParserRULE_reducer      = 2
	ReduxParserRULE_method       = 3
	ReduxParserRULE_data         = 4
	ReduxParserRULE_typeFlow     = 5
	ReduxParserRULE_define       = 6
	ReduxParserRULE_defineGlobal = 7
	ReduxParserRULE_expr         = 8
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ReduxParserT__0)|(1<<ReduxParserT__11)|(1<<ReduxParserT__12)|(1<<ReduxParserT__13))) != 0) {
		{
			p.SetState(18)
			p.Stat()
		}

		p.SetState(21)
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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Reducer()
		}

	case ReduxParserT__11, ReduxParserT__12, ReduxParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
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

	// IsReducerContext differentiates from other interfaces.
	IsReducerContext()
}

type ReducerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ReducerContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(ReduxParserIdent)
}

func (s *ReducerContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(ReduxParserIdent, i)
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
		p.SetState(27)
		p.Match(ReduxParserT__0)
	}
	{
		p.SetState(28)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(29)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(30)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(31)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(32)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(33)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(34)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(35)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(36)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(37)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(38)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(39)
		p.Match(ReduxParserT__7)
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
	p.EnterRule(localctx, 6, ReduxParserRULE_method)
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
		p.SetState(41)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(42)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(43)
		p.TypeFlow()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ReduxParserT__3 {
		{
			p.SetState(44)
			p.Match(ReduxParserT__3)
		}
		{
			p.SetState(45)
			p.TypeFlow()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(52)
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
	p.EnterRule(localctx, 8, ReduxParserRULE_data)

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
		p.SetState(54)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(55)
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
	p.EnterRule(localctx, 10, ReduxParserRULE_typeFlow)
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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(ReduxParserIdent)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(ReduxParserT__8)
		}
		{
			p.SetState(59)
			p.Match(ReduxParserT__6)
		}
		{
			p.SetState(60)
			p.Data()
		}
		{
			p.SetState(61)
			p.Match(ReduxParserT__7)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Match(ReduxParserT__9)
		}
		{
			p.SetState(64)
			p.Match(ReduxParserT__6)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ReduxParserIdent {
			{
				p.SetState(65)
				p.Method()
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(ReduxParserT__7)
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
	p.EnterRule(localctx, 12, ReduxParserRULE_define)

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
		p.SetState(75)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(76)
		p.TypeFlow()
	}
	{
		p.SetState(77)
		p.Match(ReduxParserT__10)
	}
	{
		p.SetState(78)
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
	p.EnterRule(localctx, 14, ReduxParserRULE_defineGlobal)

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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(ReduxParserT__11)
		}
		{
			p.SetState(81)
			p.Define()
		}

	case ReduxParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(ReduxParserT__12)
		}
		{
			p.SetState(83)
			p.Define()
		}

	case ReduxParserT__13:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(ReduxParserT__13)
		}
		{
			p.SetState(85)
			p.Match(ReduxParserIdent)
		}
		{
			p.SetState(86)
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ReduxParserRULE_expr, _p)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__1:
		{
			p.SetState(90)
			p.Match(ReduxParserT__1)
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		{
			p.SetState(92)
			p.Match(ReduxParserT__5)
		}

	case ReduxParserNumber:
		{
			p.SetState(94)
			p.Match(ReduxParserNumber)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ReduxParserRULE_expr)
			p.SetState(97)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			p.SetState(98)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ReduxParserT__14)|(1<<ReduxParserT__15)|(1<<ReduxParserT__16)|(1<<ReduxParserT__17))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(99)
				p.expr(4)
			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ReduxParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
