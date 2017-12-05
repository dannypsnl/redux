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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 30, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 6, 2, 10, 10, 2, 13, 2, 14, 2, 11,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 27, 2, 9, 3, 2,
	2, 2, 4, 13, 3, 2, 2, 2, 6, 15, 3, 2, 2, 2, 8, 10, 5, 4, 3, 2, 9, 8, 3,
	2, 2, 2, 10, 11, 3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12,
	3, 3, 2, 2, 2, 13, 14, 5, 6, 4, 2, 14, 5, 3, 2, 2, 2, 15, 16, 7, 3, 2,
	2, 16, 17, 7, 12, 2, 2, 17, 18, 7, 4, 2, 2, 18, 19, 7, 12, 2, 2, 19, 20,
	7, 5, 2, 2, 20, 21, 7, 6, 2, 2, 21, 22, 7, 12, 2, 2, 22, 23, 7, 7, 2, 2,
	23, 24, 7, 8, 2, 2, 24, 25, 7, 5, 2, 2, 25, 26, 7, 9, 2, 2, 26, 27, 7,
	12, 2, 2, 27, 28, 7, 10, 2, 2, 28, 7, 3, 2, 2, 2, 3, 11,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "'interface{}'", "','", "'Action'", "')'", "'{'",
	"'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "WS", "Ident",
}

var ruleNames = []string{
	"prog", "stat", "reducer",
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
	ReduxParserEOF   = antlr.TokenEOF
	ReduxParserT__0  = 1
	ReduxParserT__1  = 2
	ReduxParserT__2  = 3
	ReduxParserT__3  = 4
	ReduxParserT__4  = 5
	ReduxParserT__5  = 6
	ReduxParserT__6  = 7
	ReduxParserT__7  = 8
	ReduxParserWS    = 9
	ReduxParserIdent = 10
)

// ReduxParser rules.
const (
	ReduxParserRULE_prog    = 0
	ReduxParserRULE_stat    = 1
	ReduxParserRULE_reducer = 2
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
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ReduxParserT__0 {
		{
			p.SetState(6)
			p.Stat()
		}

		p.SetState(9)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(11)
		p.Reducer()
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
		p.SetState(13)
		p.Match(ReduxParserT__0)
	}
	{
		p.SetState(14)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(15)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(16)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(17)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(18)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(19)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(20)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(21)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(22)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(23)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(24)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(25)
		p.Match(ReduxParserT__7)
	}

	return localctx
}
