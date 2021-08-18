// Code generated from /Users/rafaeldurelli/Documents/development/golang/parserLogTest/Logging.g4 by ANTLR 4.9.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 126, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 
	4, 3, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 
	5, 5, 5, 60, 10, 5, 3, 6, 3, 6, 7, 6, 64, 10, 6, 12, 6, 14, 6, 67, 11, 
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 
	11, 90, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 96, 10, 13, 3, 14, 3, 
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 
	17, 6, 17, 121, 10, 17, 13, 17, 14, 17, 122, 3, 17, 3, 17, 2, 2, 18, 3, 
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 3, 2, 7, 6, 2, 38, 38, 67, 92, 
	97, 97, 99, 124, 7, 2, 38, 38, 50, 59, 67, 92, 97, 97, 99, 124, 8, 2, 36, 
	37, 41, 41, 46, 46, 49, 49, 60, 61, 66, 66, 3, 2, 50, 59, 5, 2, 11, 12, 
	14, 15, 34, 34, 2, 133, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 
	7, 51, 3, 2, 2, 2, 9, 59, 3, 2, 2, 2, 11, 61, 3, 2, 2, 2, 13, 68, 3, 2, 
	2, 2, 15, 70, 3, 2, 2, 2, 17, 73, 3, 2, 2, 2, 19, 75, 3, 2, 2, 2, 21, 89, 
	3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 95, 3, 2, 2, 2, 27, 97, 3, 2, 2, 2, 
	29, 108, 3, 2, 2, 2, 31, 117, 3, 2, 2, 2, 33, 120, 3, 2, 2, 2, 35, 36, 
	7, 93, 2, 2, 36, 37, 7, 73, 2, 2, 37, 38, 7, 75, 2, 2, 38, 39, 7, 80, 2, 
	2, 39, 40, 7, 95, 2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 48, 2, 2, 42, 6, 
	3, 2, 2, 2, 43, 44, 7, 82, 2, 2, 44, 45, 7, 81, 2, 2, 45, 46, 7, 85, 2, 
	2, 46, 52, 7, 86, 2, 2, 47, 48, 7, 114, 2, 2, 48, 49, 7, 113, 2, 2, 49, 
	50, 7, 117, 2, 2, 50, 52, 7, 118, 2, 2, 51, 43, 3, 2, 2, 2, 51, 47, 3, 
	2, 2, 2, 52, 8, 3, 2, 2, 2, 53, 54, 7, 73, 2, 2, 54, 55, 7, 71, 2, 2, 55, 
	60, 7, 86, 2, 2, 56, 57, 7, 105, 2, 2, 57, 58, 7, 103, 2, 2, 58, 60, 7, 
	118, 2, 2, 59, 53, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 60, 10, 3, 2, 2, 2, 
	61, 65, 5, 13, 7, 2, 62, 64, 5, 17, 9, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3, 
	2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 12, 3, 2, 2, 2, 67, 
	65, 3, 2, 2, 2, 68, 69, 9, 2, 2, 2, 69, 14, 3, 2, 2, 2, 70, 71, 7, 183, 
	2, 2, 71, 72, 7, 117, 2, 2, 72, 16, 3, 2, 2, 2, 73, 74, 9, 3, 2, 2, 74, 
	18, 3, 2, 2, 2, 75, 76, 7, 126, 2, 2, 76, 20, 3, 2, 2, 2, 77, 78, 7, 52, 
	2, 2, 78, 79, 7, 50, 2, 2, 79, 90, 7, 50, 2, 2, 80, 81, 7, 54, 2, 2, 81, 
	82, 7, 50, 2, 2, 82, 90, 7, 54, 2, 2, 83, 84, 7, 52, 2, 2, 84, 85, 7, 50, 
	2, 2, 85, 90, 7, 51, 2, 2, 86, 87, 7, 54, 2, 2, 87, 88, 7, 50, 2, 2, 88, 
	90, 7, 50, 2, 2, 89, 77, 3, 2, 2, 2, 89, 80, 3, 2, 2, 2, 89, 83, 3, 2, 
	2, 2, 89, 86, 3, 2, 2, 2, 90, 22, 3, 2, 2, 2, 91, 92, 7, 47, 2, 2, 92, 
	24, 3, 2, 2, 2, 93, 96, 9, 4, 2, 2, 94, 96, 5, 23, 12, 2, 95, 93, 3, 2, 
	2, 2, 95, 94, 3, 2, 2, 2, 96, 26, 3, 2, 2, 2, 97, 98, 5, 31, 16, 2, 98, 
	99, 5, 31, 16, 2, 99, 100, 5, 31, 16, 2, 100, 101, 5, 31, 16, 2, 101, 102, 
	7, 49, 2, 2, 102, 103, 5, 31, 16, 2, 103, 104, 5, 31, 16, 2, 104, 105, 
	7, 49, 2, 2, 105, 106, 5, 31, 16, 2, 106, 107, 5, 31, 16, 2, 107, 28, 3, 
	2, 2, 2, 108, 109, 5, 31, 16, 2, 109, 110, 5, 31, 16, 2, 110, 111, 7, 60, 
	2, 2, 111, 112, 5, 31, 16, 2, 112, 113, 5, 31, 16, 2, 113, 114, 7, 60, 
	2, 2, 114, 115, 5, 31, 16, 2, 115, 116, 5, 31, 16, 2, 116, 30, 3, 2, 2, 
	2, 117, 118, 9, 5, 2, 2, 118, 32, 3, 2, 2, 2, 119, 121, 9, 6, 2, 2, 120, 
	119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 
	3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 8, 17, 2, 2, 125, 34, 3, 2, 
	2, 2, 9, 2, 51, 59, 65, 89, 95, 122, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'[GIN]'", "'.'", "", "", "", "", "'\u00B5s'", "", "'|'", "", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "POST", "GET", "Identifier", "Letter", "Mu", "LetterOrDigit", 
	"PIPE", "CODE", "HYPHEN", "SPECIAL", "DATE", "TMSTAMP", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "POST", "GET", "Identifier", "Letter", "Mu", "LetterOrDigit", 
	"PIPE", "CODE", "HYPHEN", "SPECIAL", "DATE", "TMSTAMP", "DIGIT", "WS",
}
type LoggingLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewLoggingLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *LoggingLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLoggingLexer(input antlr.CharStream) *LoggingLexer {
	l := new(LoggingLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Logging.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LoggingLexer tokens.
const (
	LoggingLexerT__0 = 1
	LoggingLexerT__1 = 2
	LoggingLexerPOST = 3
	LoggingLexerGET = 4
	LoggingLexerIdentifier = 5
	LoggingLexerLetter = 6
	LoggingLexerMu = 7
	LoggingLexerLetterOrDigit = 8
	LoggingLexerPIPE = 9
	LoggingLexerCODE = 10
	LoggingLexerHYPHEN = 11
	LoggingLexerSPECIAL = 12
	LoggingLexerDATE = 13
	LoggingLexerTMSTAMP = 14
	LoggingLexerDIGIT = 15
	LoggingLexerWS = 16
)

