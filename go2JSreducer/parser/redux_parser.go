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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 19, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 17, 2, 4, 3, 2, 2, 2, 4,
	5, 7, 3, 2, 2, 5, 6, 7, 12, 2, 2, 6, 7, 7, 4, 2, 2, 7, 8, 7, 12, 2, 2,
	8, 9, 7, 5, 2, 2, 9, 10, 7, 6, 2, 2, 10, 11, 7, 12, 2, 2, 11, 12, 7, 7,
	2, 2, 12, 13, 7, 8, 2, 2, 13, 14, 7, 5, 2, 2, 14, 15, 7, 9, 2, 2, 15, 16,
	7, 12, 2, 2, 16, 17, 7, 10, 2, 2, 17, 3, 3, 2, 2, 2, 2,
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
	"reducer",
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

// ReduxParserRULE_reducer is the ReduxParser rule.
const ReduxParserRULE_reducer = 0

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
	p.EnterRule(localctx, 0, ReduxParserRULE_reducer)

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
		p.SetState(2)
		p.Match(ReduxParserT__0)
	}
	{
		p.SetState(3)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(4)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(5)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(6)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(7)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(8)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(9)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(10)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(11)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(12)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(13)
		p.Match(ReduxParserIdent)
	}
	{
		p.SetState(14)
		p.Match(ReduxParserT__7)
	}

	return localctx
}
