// Generated from Redux.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 86, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 6, 10, 65, 10, 10, 13, 10, 14, 10, 66, 3, 10, 3, 10,
	3, 11, 3, 11, 7, 11, 73, 10, 11, 12, 11, 14, 11, 76, 11, 11, 3, 12, 5,
	12, 79, 10, 12, 3, 13, 3, 13, 5, 13, 83, 10, 13, 3, 14, 3, 14, 2, 2, 15,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	2, 25, 2, 27, 2, 3, 2, 6, 5, 2, 11, 12, 15, 15, 34, 34, 16, 2, 67, 92,
	97, 97, 99, 124, 194, 216, 218, 248, 250, 769, 882, 895, 897, 8193, 8206,
	8207, 8306, 8593, 11266, 12273, 12291, 55297, 63746, 64977, 65010, 65535,
	6, 2, 50, 59, 185, 185, 770, 881, 8257, 8258, 3, 2, 50, 59, 2, 85, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 34, 3, 2, 2, 2,
	7, 36, 3, 2, 2, 2, 9, 48, 3, 2, 2, 2, 11, 50, 3, 2, 2, 2, 13, 57, 3, 2,
	2, 2, 15, 59, 3, 2, 2, 2, 17, 61, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 70,
	3, 2, 2, 2, 23, 78, 3, 2, 2, 2, 25, 82, 3, 2, 2, 2, 27, 84, 3, 2, 2, 2,
	29, 30, 7, 104, 2, 2, 30, 31, 7, 119, 2, 2, 31, 32, 7, 112, 2, 2, 32, 33,
	7, 101, 2, 2, 33, 4, 3, 2, 2, 2, 34, 35, 7, 42, 2, 2, 35, 6, 3, 2, 2, 2,
	36, 37, 7, 107, 2, 2, 37, 38, 7, 112, 2, 2, 38, 39, 7, 118, 2, 2, 39, 40,
	7, 103, 2, 2, 40, 41, 7, 116, 2, 2, 41, 42, 7, 104, 2, 2, 42, 43, 7, 99,
	2, 2, 43, 44, 7, 101, 2, 2, 44, 45, 7, 103, 2, 2, 45, 46, 7, 125, 2, 2,
	46, 47, 7, 127, 2, 2, 47, 8, 3, 2, 2, 2, 48, 49, 7, 46, 2, 2, 49, 10, 3,
	2, 2, 2, 50, 51, 7, 67, 2, 2, 51, 52, 7, 101, 2, 2, 52, 53, 7, 118, 2,
	2, 53, 54, 7, 107, 2, 2, 54, 55, 7, 113, 2, 2, 55, 56, 7, 112, 2, 2, 56,
	12, 3, 2, 2, 2, 57, 58, 7, 43, 2, 2, 58, 14, 3, 2, 2, 2, 59, 60, 7, 125,
	2, 2, 60, 16, 3, 2, 2, 2, 61, 62, 7, 127, 2, 2, 62, 18, 3, 2, 2, 2, 63,
	65, 9, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2,
	2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 8, 10, 2, 2, 69, 20,
	3, 2, 2, 2, 70, 74, 5, 23, 12, 2, 71, 73, 5, 25, 13, 2, 72, 71, 3, 2, 2,
	2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 22,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79, 9, 3, 2, 2, 78, 77, 3, 2, 2, 2,
	79, 24, 3, 2, 2, 2, 80, 83, 5, 23, 12, 2, 81, 83, 9, 4, 2, 2, 82, 80, 3,
	2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 26, 3, 2, 2, 2, 84, 85, 9, 5, 2, 2, 85,
	28, 3, 2, 2, 2, 7, 2, 66, 74, 78, 82, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'func'", "'('", "'interface{}'", "','", "'Action'", "')'", "'{'",
	"'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "WS", "Ident",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "WS", "Ident",
	"StartLetter", "Letter", "Digit",
}

type ReduxLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewReduxLexer(input antlr.CharStream) *ReduxLexer {

	l := new(ReduxLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Redux.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ReduxLexer tokens.
const (
	ReduxLexerT__0  = 1
	ReduxLexerT__1  = 2
	ReduxLexerT__2  = 3
	ReduxLexerT__3  = 4
	ReduxLexerT__4  = 5
	ReduxLexerT__5  = 6
	ReduxLexerT__6  = 7
	ReduxLexerT__7  = 8
	ReduxLexerWS    = 9
	ReduxLexerIdent = 10
)
