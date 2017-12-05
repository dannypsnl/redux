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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 190,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 30, 10, 2, 13, 2, 14, 2, 31, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 49, 10, 4, 12, 4, 14, 4, 52, 11, 4, 5, 4, 54, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 99, 10, 6, 13, 6, 14,
	6, 100, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 108, 10, 7, 12, 7, 14, 7, 111,
	11, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 128, 10, 9, 12, 9, 14, 9, 131, 11, 9, 3,
	9, 3, 9, 5, 9, 135, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 149, 10, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 159, 10, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 169, 10, 12, 12, 12,
	14, 12, 172, 11, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 180,
	10, 14, 12, 14, 14, 14, 183, 11, 14, 3, 14, 3, 14, 3, 14, 5, 14, 188, 10,
	14, 3, 14, 2, 3, 22, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	2, 3, 3, 2, 27, 30, 2, 198, 2, 29, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 40,
	3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 98, 3, 2, 2, 2, 12, 102, 3, 2, 2, 2,
	14, 115, 3, 2, 2, 2, 16, 134, 3, 2, 2, 2, 18, 136, 3, 2, 2, 2, 20, 148,
	3, 2, 2, 2, 22, 158, 3, 2, 2, 2, 24, 173, 3, 2, 2, 2, 26, 187, 3, 2, 2,
	2, 28, 30, 5, 4, 3, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 3, 3, 2, 2, 2, 33, 39, 5, 8, 5, 2,
	34, 39, 5, 6, 4, 2, 35, 39, 5, 20, 11, 2, 36, 39, 5, 24, 13, 2, 37, 39,
	5, 26, 14, 2, 38, 33, 3, 2, 2, 2, 38, 34, 3, 2, 2, 2, 38, 35, 3, 2, 2,
	2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 5, 3, 2, 2, 2, 40, 41, 7,
	3, 2, 2, 41, 42, 7, 35, 2, 2, 42, 53, 7, 4, 2, 2, 43, 44, 7, 35, 2, 2,
	44, 50, 5, 16, 9, 2, 45, 46, 7, 5, 2, 2, 46, 47, 7, 35, 2, 2, 47, 49, 5,
	16, 9, 2, 48, 45, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 43, 3, 2, 2,
	2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 7, 6, 2, 2, 56, 57,
	7, 7, 2, 2, 57, 58, 5, 4, 3, 2, 58, 59, 7, 8, 2, 2, 59, 7, 3, 2, 2, 2,
	60, 61, 7, 3, 2, 2, 61, 62, 7, 35, 2, 2, 62, 63, 7, 4, 2, 2, 63, 64, 7,
	35, 2, 2, 64, 65, 7, 9, 2, 2, 65, 66, 7, 5, 2, 2, 66, 67, 7, 35, 2, 2,
	67, 68, 7, 10, 2, 2, 68, 69, 7, 6, 2, 2, 69, 70, 7, 9, 2, 2, 70, 71, 7,
	7, 2, 2, 71, 72, 7, 11, 2, 2, 72, 73, 7, 35, 2, 2, 73, 74, 7, 12, 2, 2,
	74, 75, 7, 13, 2, 2, 75, 76, 7, 7, 2, 2, 76, 77, 7, 14, 2, 2, 77, 78, 5,
	22, 12, 2, 78, 79, 7, 8, 2, 2, 79, 80, 7, 15, 2, 2, 80, 81, 7, 35, 2, 2,
	81, 82, 7, 16, 2, 2, 82, 83, 7, 17, 2, 2, 83, 84, 7, 7, 2, 2, 84, 85, 5,
	10, 6, 2, 85, 86, 7, 18, 2, 2, 86, 87, 7, 19, 2, 2, 87, 88, 7, 14, 2, 2,
	88, 89, 5, 22, 12, 2, 89, 90, 7, 8, 2, 2, 90, 91, 7, 8, 2, 2, 91, 9, 3,
	2, 2, 2, 92, 93, 7, 20, 2, 2, 93, 94, 5, 22, 12, 2, 94, 95, 7, 19, 2, 2,
	95, 96, 7, 14, 2, 2, 96, 97, 5, 22, 12, 2, 97, 99, 3, 2, 2, 2, 98, 92,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 11, 3, 2, 2, 2, 102, 103, 7, 35, 2, 2, 103, 104, 7, 4, 2, 2, 104,
	109, 5, 16, 9, 2, 105, 106, 7, 5, 2, 2, 106, 108, 5, 16, 9, 2, 107, 105,
	3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2,
	2, 2, 110, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 113, 7, 6, 2, 2,
	113, 114, 5, 16, 9, 2, 114, 13, 3, 2, 2, 2, 115, 116, 7, 35, 2, 2, 116,
	117, 5, 16, 9, 2, 117, 15, 3, 2, 2, 2, 118, 135, 7, 35, 2, 2, 119, 120,
	7, 21, 2, 2, 120, 121, 7, 7, 2, 2, 121, 122, 5, 14, 8, 2, 122, 123, 7,
	8, 2, 2, 123, 135, 3, 2, 2, 2, 124, 125, 7, 22, 2, 2, 125, 129, 7, 7, 2,
	2, 126, 128, 5, 12, 7, 2, 127, 126, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129,
	127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 132, 135, 7, 8, 2, 2, 133, 135, 3, 2, 2, 2, 134, 118, 3, 2,
	2, 2, 134, 119, 3, 2, 2, 2, 134, 124, 3, 2, 2, 2, 134, 133, 3, 2, 2, 2,
	135, 17, 3, 2, 2, 2, 136, 137, 7, 35, 2, 2, 137, 138, 5, 16, 9, 2, 138,
	139, 7, 23, 2, 2, 139, 140, 5, 22, 12, 2, 140, 19, 3, 2, 2, 2, 141, 142,
	7, 24, 2, 2, 142, 149, 5, 18, 10, 2, 143, 144, 7, 25, 2, 2, 144, 149, 5,
	18, 10, 2, 145, 146, 7, 26, 2, 2, 146, 147, 7, 35, 2, 2, 147, 149, 5, 16,
	9, 2, 148, 141, 3, 2, 2, 2, 148, 143, 3, 2, 2, 2, 148, 145, 3, 2, 2, 2,
	149, 21, 3, 2, 2, 2, 150, 151, 8, 12, 1, 2, 151, 152, 7, 4, 2, 2, 152,
	153, 5, 22, 12, 2, 153, 154, 7, 6, 2, 2, 154, 159, 3, 2, 2, 2, 155, 159,
	7, 36, 2, 2, 156, 159, 7, 34, 2, 2, 157, 159, 7, 35, 2, 2, 158, 150, 3,
	2, 2, 2, 158, 155, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 157, 3, 2, 2,
	2, 159, 170, 3, 2, 2, 2, 160, 161, 12, 8, 2, 2, 161, 162, 9, 2, 2, 2, 162,
	169, 5, 22, 12, 9, 163, 164, 12, 6, 2, 2, 164, 165, 7, 16, 2, 2, 165, 166,
	7, 4, 2, 2, 166, 167, 7, 35, 2, 2, 167, 169, 7, 6, 2, 2, 168, 160, 3, 2,
	2, 2, 168, 163, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2,
	170, 171, 3, 2, 2, 2, 171, 23, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 174,
	7, 31, 2, 2, 174, 175, 7, 35, 2, 2, 175, 25, 3, 2, 2, 2, 176, 177, 7, 32,
	2, 2, 177, 181, 7, 4, 2, 2, 178, 180, 7, 34, 2, 2, 179, 178, 3, 2, 2, 2,
	180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182,
	184, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 188, 7, 6, 2, 2, 185, 186,
	7, 32, 2, 2, 186, 188, 7, 34, 2, 2, 187, 176, 3, 2, 2, 2, 187, 185, 3,
	2, 2, 2, 188, 27, 3, 2, 2, 2, 16, 31, 38, 50, 53, 100, 109, 129, 134, 148,
	158, 168, 170, 181, 187,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "','", "')'", "'{'", "'}'", "'interface{}'", "'redux.Action'",
	"'if'", "'=='", "'nil'", "'return'", "'switch'", "'.'", "'Type'", "'default'",
	"':'", "'case'", "'struct'", "'interface'", "'='", "'const'", "'var'",
	"'type'", "'+'", "'-'", "'*'", "'/'", "'package'", "'import'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "STRING", "ID",
	"Number",
}

var ruleNames = []string{
	"prog", "stat", "function", "reducer", "cases", "method", "data", "typeFlow",
	"define", "defineGlobal", "expr", "packageStat", "importStat",
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
	ReduxParserT__28  = 29
	ReduxParserT__29  = 30
	ReduxParserWS     = 31
	ReduxParserSTRING = 32
	ReduxParserID     = 33
	ReduxParserNumber = 34
)

// ReduxParser rules.
const (
	ReduxParserRULE_prog         = 0
	ReduxParserRULE_stat         = 1
	ReduxParserRULE_function     = 2
	ReduxParserRULE_reducer      = 3
	ReduxParserRULE_cases        = 4
	ReduxParserRULE_method       = 5
	ReduxParserRULE_data         = 6
	ReduxParserRULE_typeFlow     = 7
	ReduxParserRULE_define       = 8
	ReduxParserRULE_defineGlobal = 9
	ReduxParserRULE_expr         = 10
	ReduxParserRULE_packageStat  = 11
	ReduxParserRULE_importStat   = 12
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ReduxParserT__0)|(1<<ReduxParserT__21)|(1<<ReduxParserT__22)|(1<<ReduxParserT__23)|(1<<ReduxParserT__28)|(1<<ReduxParserT__29))) != 0) {
		{
			p.SetState(26)
			p.Stat()
		}

		p.SetState(29)
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

func (s *StatContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *StatContext) DefineGlobal() IDefineGlobalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineGlobalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineGlobalContext)
}

func (s *StatContext) PackageStat() IPackageStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageStatContext)
}

func (s *StatContext) ImportStat() IImportStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportStatContext)
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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Reducer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Function()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.DefineGlobal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(34)
			p.PackageStat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(35)
			p.ImportStat()
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ReduxParserID)
}

func (s *FunctionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ReduxParserID, i)
}

func (s *FunctionContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *FunctionContext) AllTypeFlow() []ITypeFlowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem())
	var tst = make([]ITypeFlowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFlowContext)
		}
	}

	return tst
}

func (s *FunctionContext) TypeFlow(i int) ITypeFlowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFlowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFlowContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *ReduxParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ReduxParserRULE_function)
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
		p.SetState(38)
		p.Match(ReduxParserT__0)
	}
	{
		p.SetState(39)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(40)
		p.Match(ReduxParserT__1)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ReduxParserID {
		{
			p.SetState(41)
			p.Match(ReduxParserID)
		}
		{
			p.SetState(42)
			p.TypeFlow()
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ReduxParserT__2 {
			{
				p.SetState(43)
				p.Match(ReduxParserT__2)
			}
			{
				p.SetState(44)
				p.Match(ReduxParserID)
			}
			{
				p.SetState(45)
				p.TypeFlow()
			}

			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(53)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(54)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(55)
		p.Stat()
	}
	{
		p.SetState(56)
		p.Match(ReduxParserT__5)
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

func (s *ReducerContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ReduxParserID)
}

func (s *ReducerContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ReduxParserID, i)
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
	p.EnterRule(localctx, 6, ReduxParserRULE_reducer)

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
		p.SetState(58)
		p.Match(ReduxParserT__0)
	}
	{
		p.SetState(59)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(60)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(61)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(62)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(63)
		p.Match(ReduxParserT__2)
	}
	{
		p.SetState(64)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(65)
		p.Match(ReduxParserT__7)
	}
	{
		p.SetState(66)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(67)
		p.Match(ReduxParserT__6)
	}
	{
		p.SetState(68)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(69)
		p.Match(ReduxParserT__8)
	}
	{
		p.SetState(70)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(71)
		p.Match(ReduxParserT__9)
	}
	{
		p.SetState(72)
		p.Match(ReduxParserT__10)
	}
	{
		p.SetState(73)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(74)
		p.Match(ReduxParserT__11)
	}
	{
		p.SetState(75)

		var _x = p.expr(0)

		localctx.(*ReducerContext).initialState = _x
	}
	{
		p.SetState(76)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(77)
		p.Match(ReduxParserT__12)
	}
	{
		p.SetState(78)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(79)
		p.Match(ReduxParserT__13)
	}
	{
		p.SetState(80)
		p.Match(ReduxParserT__14)
	}
	{
		p.SetState(81)
		p.Match(ReduxParserT__4)
	}
	{
		p.SetState(82)
		p.Cases()
	}
	{
		p.SetState(83)
		p.Match(ReduxParserT__15)
	}
	{
		p.SetState(84)
		p.Match(ReduxParserT__16)
	}
	{
		p.SetState(85)
		p.Match(ReduxParserT__11)
	}
	{
		p.SetState(86)
		p.expr(0)
	}
	{
		p.SetState(87)
		p.Match(ReduxParserT__5)
	}
	{
		p.SetState(88)
		p.Match(ReduxParserT__5)
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
	p.EnterRule(localctx, 8, ReduxParserRULE_cases)
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
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ReduxParserT__17 {
		{
			p.SetState(90)
			p.Match(ReduxParserT__17)
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		{
			p.SetState(92)
			p.Match(ReduxParserT__16)
		}
		{
			p.SetState(93)
			p.Match(ReduxParserT__11)
		}
		{
			p.SetState(94)
			p.expr(0)
		}

		p.SetState(98)
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

func (s *MethodContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
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
	p.EnterRule(localctx, 10, ReduxParserRULE_method)
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
		p.SetState(100)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(101)
		p.Match(ReduxParserT__1)
	}
	{
		p.SetState(102)
		p.TypeFlow()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ReduxParserT__2 {
		{
			p.SetState(103)
			p.Match(ReduxParserT__2)
		}
		{
			p.SetState(104)
			p.TypeFlow()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(ReduxParserT__3)
	}
	{
		p.SetState(111)
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

func (s *DataContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
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
	p.EnterRule(localctx, 12, ReduxParserRULE_data)

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
		p.SetState(113)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(114)
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

func (s *TypeFlowContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
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
	p.EnterRule(localctx, 14, ReduxParserRULE_typeFlow)
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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(ReduxParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(ReduxParserT__18)
		}
		{
			p.SetState(118)
			p.Match(ReduxParserT__4)
		}
		{
			p.SetState(119)
			p.Data()
		}
		{
			p.SetState(120)
			p.Match(ReduxParserT__5)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.Match(ReduxParserT__19)
		}
		{
			p.SetState(123)
			p.Match(ReduxParserT__4)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ReduxParserID {
			{
				p.SetState(124)
				p.Method()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(130)
			p.Match(ReduxParserT__5)
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

func (s *DefineContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
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
	p.EnterRule(localctx, 16, ReduxParserRULE_define)

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
		p.SetState(134)
		p.Match(ReduxParserID)
	}
	{
		p.SetState(135)
		p.TypeFlow()
	}
	{
		p.SetState(136)
		p.Match(ReduxParserT__20)
	}
	{
		p.SetState(137)
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

func (s *DefineGlobalContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
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
	p.EnterRule(localctx, 18, ReduxParserRULE_defineGlobal)

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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Match(ReduxParserT__21)
		}
		{
			p.SetState(140)
			p.Define()
		}

	case ReduxParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(ReduxParserT__22)
		}
		{
			p.SetState(142)
			p.Define()
		}

	case ReduxParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)
			p.Match(ReduxParserT__23)
		}
		{
			p.SetState(144)
			p.Match(ReduxParserID)
		}
		{
			p.SetState(145)
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

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ReduxParserSTRING, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, ReduxParserRULE_expr, _p)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ReduxParserT__1:
		{
			p.SetState(149)
			p.Match(ReduxParserT__1)
		}
		{
			p.SetState(150)
			p.expr(0)
		}
		{
			p.SetState(151)
			p.Match(ReduxParserT__3)
		}

	case ReduxParserNumber:
		{
			p.SetState(153)
			p.Match(ReduxParserNumber)
		}

	case ReduxParserSTRING:
		{
			p.SetState(154)
			p.Match(ReduxParserSTRING)
		}

	case ReduxParserID:
		{
			p.SetState(155)
			p.Match(ReduxParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ReduxParserRULE_expr)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(159)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ReduxParserT__24)|(1<<ReduxParserT__25)|(1<<ReduxParserT__26)|(1<<ReduxParserT__27))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(160)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ReduxParserRULE_expr)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(162)
					p.Match(ReduxParserT__13)
				}
				{
					p.SetState(163)
					p.Match(ReduxParserT__1)
				}
				{
					p.SetState(164)
					p.Match(ReduxParserID)
				}
				{
					p.SetState(165)
					p.Match(ReduxParserT__3)
				}

			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IPackageStatContext is an interface to support dynamic dispatch.
type IPackageStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageStatContext differentiates from other interfaces.
	IsPackageStatContext()
}

type PackageStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatContext() *PackageStatContext {
	var p = new(PackageStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_packageStat
	return p
}

func (*PackageStatContext) IsPackageStatContext() {}

func NewPackageStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatContext {
	var p = new(PackageStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_packageStat

	return p
}

func (s *PackageStatContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatContext) ID() antlr.TerminalNode {
	return s.GetToken(ReduxParserID, 0)
}

func (s *PackageStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterPackageStat(s)
	}
}

func (s *PackageStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitPackageStat(s)
	}
}

func (p *ReduxParser) PackageStat() (localctx IPackageStatContext) {
	localctx = NewPackageStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ReduxParserRULE_packageStat)

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
		p.SetState(171)
		p.Match(ReduxParserT__28)
	}
	{
		p.SetState(172)
		p.Match(ReduxParserID)
	}

	return localctx
}

// IImportStatContext is an interface to support dynamic dispatch.
type IImportStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatContext differentiates from other interfaces.
	IsImportStatContext()
}

type ImportStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatContext() *ImportStatContext {
	var p = new(ImportStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReduxParserRULE_importStat
	return p
}

func (*ImportStatContext) IsImportStatContext() {}

func NewImportStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatContext {
	var p = new(ImportStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReduxParserRULE_importStat

	return p
}

func (s *ImportStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ReduxParserSTRING)
}

func (s *ImportStatContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ReduxParserSTRING, i)
}

func (s *ImportStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.EnterImportStat(s)
	}
}

func (s *ImportStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReduxListener); ok {
		listenerT.ExitImportStat(s)
	}
}

func (p *ReduxParser) ImportStat() (localctx IImportStatContext) {
	localctx = NewImportStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ReduxParserRULE_importStat)
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

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.Match(ReduxParserT__29)
		}
		{
			p.SetState(175)
			p.Match(ReduxParserT__1)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ReduxParserSTRING {
			{
				p.SetState(176)
				p.Match(ReduxParserSTRING)
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(182)
			p.Match(ReduxParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(ReduxParserT__29)
		}
		{
			p.SetState(184)
			p.Match(ReduxParserSTRING)
		}

	}

	return localctx
}

func (p *ReduxParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
