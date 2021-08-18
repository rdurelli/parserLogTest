// Code generated from /Users/rafaeldurelli/Documents/development/golang/parserLogTest/Logging.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Logging

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 91, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 7, 
	2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 3, 2, 3, 3, 7, 3, 34, 10, 
	3, 12, 3, 14, 3, 37, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 7, 6, 55, 10, 6, 12, 
	6, 14, 6, 58, 11, 6, 3, 6, 6, 6, 61, 10, 6, 13, 6, 14, 6, 62, 3, 6, 7, 
	6, 66, 10, 6, 12, 6, 14, 6, 69, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 5, 7, 75, 
	10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 7, 10, 82, 10, 10, 12, 10, 14, 10, 
	85, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 83, 2, 13, 2, 4, 6, 8, 
	10, 12, 14, 16, 18, 20, 22, 2, 2, 2, 86, 2, 27, 3, 2, 2, 2, 4, 35, 3, 2, 
	2, 2, 6, 49, 3, 2, 2, 2, 8, 51, 3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 74, 
	3, 2, 2, 2, 14, 76, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 
	20, 86, 3, 2, 2, 2, 22, 88, 3, 2, 2, 2, 24, 26, 5, 4, 3, 2, 25, 24, 3, 
	2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 
	30, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 31, 7, 2, 2, 3, 31, 3, 3, 2, 2, 
	2, 32, 34, 5, 6, 4, 2, 33, 32, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 
	3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 
	38, 39, 5, 8, 5, 2, 39, 40, 7, 13, 2, 2, 40, 41, 7, 16, 2, 2, 41, 42, 5, 
	20, 11, 2, 42, 43, 5, 22, 12, 2, 43, 44, 5, 20, 11, 2, 44, 45, 5, 10, 6, 
	2, 45, 46, 5, 20, 11, 2, 46, 47, 5, 12, 7, 2, 47, 48, 5, 18, 10, 2, 48, 
	5, 3, 2, 2, 2, 49, 50, 7, 3, 2, 2, 50, 7, 3, 2, 2, 2, 51, 52, 7, 15, 2, 
	2, 52, 9, 3, 2, 2, 2, 53, 55, 7, 10, 2, 2, 54, 53, 3, 2, 2, 2, 55, 58, 
	3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 
	58, 56, 3, 2, 2, 2, 59, 61, 7, 4, 2, 2, 60, 59, 3, 2, 2, 2, 61, 62, 3, 
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 67, 3, 2, 2, 2, 64, 
	66, 7, 10, 2, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 
	2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 71, 
	7, 9, 2, 2, 71, 11, 3, 2, 2, 2, 72, 75, 5, 14, 8, 2, 73, 75, 5, 16, 9, 
	2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 13, 3, 2, 2, 2, 76, 77, 
	7, 5, 2, 2, 77, 15, 3, 2, 2, 2, 78, 79, 7, 6, 2, 2, 79, 17, 3, 2, 2, 2, 
	80, 82, 11, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 84, 3, 
	2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 19, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 
	87, 7, 11, 2, 2, 87, 21, 3, 2, 2, 2, 88, 89, 7, 12, 2, 2, 89, 23, 3, 2, 
	2, 2, 9, 27, 35, 56, 62, 67, 74, 83,
}
var literalNames = []string{
	"", "'[GIN]'", "'.'", "", "", "", "", "'\u00B5s'", "", "'|'", "", "'-'",
}
var symbolicNames = []string{
	"", "", "", "POST", "GET", "Identifier", "Letter", "Mu", "LetterOrDigit", 
	"PIPE", "CODE", "HYPHEN", "SPECIAL", "DATE", "TMSTAMP", "DIGIT", "WS",
}

var ruleNames = []string{
	"start", "log", "gin", "dateformat", "requestTime", "httpMethod", "postMethod", 
	"getMethod", "path", "pipe", "statusCode",
}
type LoggingParser struct {
	*antlr.BaseParser
}

// NewLoggingParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LoggingParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLoggingParser(input antlr.TokenStream) *LoggingParser {
	this := new(LoggingParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Logging.g4"

	return this
}


// LoggingParser tokens.
const (
	LoggingParserEOF = antlr.TokenEOF
	LoggingParserT__0 = 1
	LoggingParserT__1 = 2
	LoggingParserPOST = 3
	LoggingParserGET = 4
	LoggingParserIdentifier = 5
	LoggingParserLetter = 6
	LoggingParserMu = 7
	LoggingParserLetterOrDigit = 8
	LoggingParserPIPE = 9
	LoggingParserCODE = 10
	LoggingParserHYPHEN = 11
	LoggingParserSPECIAL = 12
	LoggingParserDATE = 13
	LoggingParserTMSTAMP = 14
	LoggingParserDIGIT = 15
	LoggingParserWS = 16
)

// LoggingParser rules.
const (
	LoggingParserRULE_start = 0
	LoggingParserRULE_log = 1
	LoggingParserRULE_gin = 2
	LoggingParserRULE_dateformat = 3
	LoggingParserRULE_requestTime = 4
	LoggingParserRULE_httpMethod = 5
	LoggingParserRULE_postMethod = 6
	LoggingParserRULE_getMethod = 7
	LoggingParserRULE_path = 8
	LoggingParserRULE_pipe = 9
	LoggingParserRULE_statusCode = 10
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(LoggingParserEOF, 0)
}

func (s *StartContext) AllLog() []ILogContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILogContext)(nil)).Elem())
	var tst = make([]ILogContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILogContext)
		}
	}

	return tst
}

func (s *StartContext) Log(i int) ILogContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILogContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LoggingParserRULE_start)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == LoggingParserT__0 || _la == LoggingParserDATE {
		{
			p.SetState(22)
			p.Log()
		}


		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(LoggingParserEOF)
	}



	return localctx
}


// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_log
	return p
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) Dateformat() IDateformatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateformatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateformatContext)
}

func (s *LogContext) HYPHEN() antlr.TerminalNode {
	return s.GetToken(LoggingParserHYPHEN, 0)
}

func (s *LogContext) TMSTAMP() antlr.TerminalNode {
	return s.GetToken(LoggingParserTMSTAMP, 0)
}

func (s *LogContext) AllPipe() []IPipeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeContext)(nil)).Elem())
	var tst = make([]IPipeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeContext)
		}
	}

	return tst
}

func (s *LogContext) Pipe(i int) IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *LogContext) StatusCode() IStatusCodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatusCodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatusCodeContext)
}

func (s *LogContext) RequestTime() IRequestTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestTimeContext)
}

func (s *LogContext) HttpMethod() IHttpMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHttpMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHttpMethodContext)
}

func (s *LogContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *LogContext) AllGin() []IGinContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGinContext)(nil)).Elem())
	var tst = make([]IGinContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGinContext)
		}
	}

	return tst
}

func (s *LogContext) Gin(i int) IGinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGinContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGinContext)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitLog(s)
	}
}

func (s *LogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitLog(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) Log() (localctx ILogContext) {
	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LoggingParserRULE_log)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == LoggingParserT__0 {
		{
			p.SetState(30)
			p.Gin()
		}


		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Dateformat()
	}
	{
		p.SetState(37)
		p.Match(LoggingParserHYPHEN)
	}
	{
		p.SetState(38)
		p.Match(LoggingParserTMSTAMP)
	}
	{
		p.SetState(39)
		p.Pipe()
	}
	{
		p.SetState(40)
		p.StatusCode()
	}
	{
		p.SetState(41)
		p.Pipe()
	}
	{
		p.SetState(42)
		p.RequestTime()
	}
	{
		p.SetState(43)
		p.Pipe()
	}
	{
		p.SetState(44)
		p.HttpMethod()
	}
	{
		p.SetState(45)
		p.Path()
	}



	return localctx
}


// IGinContext is an interface to support dynamic dispatch.
type IGinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGinContext differentiates from other interfaces.
	IsGinContext()
}

type GinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGinContext() *GinContext {
	var p = new(GinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_gin
	return p
}

func (*GinContext) IsGinContext() {}

func NewGinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GinContext {
	var p = new(GinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_gin

	return p
}

func (s *GinContext) GetParser() antlr.Parser { return s.parser }
func (s *GinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterGin(s)
	}
}

func (s *GinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitGin(s)
	}
}

func (s *GinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitGin(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) Gin() (localctx IGinContext) {
	localctx = NewGinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LoggingParserRULE_gin)

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
		p.SetState(47)
		p.Match(LoggingParserT__0)
	}



	return localctx
}


// IDateformatContext is an interface to support dynamic dispatch.
type IDateformatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateformatContext differentiates from other interfaces.
	IsDateformatContext()
}

type DateformatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateformatContext() *DateformatContext {
	var p = new(DateformatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_dateformat
	return p
}

func (*DateformatContext) IsDateformatContext() {}

func NewDateformatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateformatContext {
	var p = new(DateformatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_dateformat

	return p
}

func (s *DateformatContext) GetParser() antlr.Parser { return s.parser }

func (s *DateformatContext) DATE() antlr.TerminalNode {
	return s.GetToken(LoggingParserDATE, 0)
}

func (s *DateformatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateformatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DateformatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterDateformat(s)
	}
}

func (s *DateformatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitDateformat(s)
	}
}

func (s *DateformatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitDateformat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) Dateformat() (localctx IDateformatContext) {
	localctx = NewDateformatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LoggingParserRULE_dateformat)

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
		p.SetState(49)
		p.Match(LoggingParserDATE)
	}



	return localctx
}


// IRequestTimeContext is an interface to support dynamic dispatch.
type IRequestTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestTimeContext differentiates from other interfaces.
	IsRequestTimeContext()
}

type RequestTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestTimeContext() *RequestTimeContext {
	var p = new(RequestTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_requestTime
	return p
}

func (*RequestTimeContext) IsRequestTimeContext() {}

func NewRequestTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestTimeContext {
	var p = new(RequestTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_requestTime

	return p
}

func (s *RequestTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestTimeContext) Mu() antlr.TerminalNode {
	return s.GetToken(LoggingParserMu, 0)
}

func (s *RequestTimeContext) AllLetterOrDigit() []antlr.TerminalNode {
	return s.GetTokens(LoggingParserLetterOrDigit)
}

func (s *RequestTimeContext) LetterOrDigit(i int) antlr.TerminalNode {
	return s.GetToken(LoggingParserLetterOrDigit, i)
}

func (s *RequestTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RequestTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterRequestTime(s)
	}
}

func (s *RequestTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitRequestTime(s)
	}
}

func (s *RequestTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitRequestTime(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) RequestTime() (localctx IRequestTimeContext) {
	localctx = NewRequestTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LoggingParserRULE_requestTime)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == LoggingParserLetterOrDigit {
		{
			p.SetState(51)
			p.Match(LoggingParserLetterOrDigit)
		}


		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LoggingParserT__1 {
		{
			p.SetState(57)
			p.Match(LoggingParserT__1)
		}


		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == LoggingParserLetterOrDigit {
		{
			p.SetState(62)
			p.Match(LoggingParserLetterOrDigit)
		}


		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(LoggingParserMu)
	}



	return localctx
}


// IHttpMethodContext is an interface to support dynamic dispatch.
type IHttpMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHttpMethodContext differentiates from other interfaces.
	IsHttpMethodContext()
}

type HttpMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpMethodContext() *HttpMethodContext {
	var p = new(HttpMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_httpMethod
	return p
}

func (*HttpMethodContext) IsHttpMethodContext() {}

func NewHttpMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpMethodContext {
	var p = new(HttpMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_httpMethod

	return p
}

func (s *HttpMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpMethodContext) PostMethod() IPostMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostMethodContext)
}

func (s *HttpMethodContext) GetMethod() IGetMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetMethodContext)
}

func (s *HttpMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *HttpMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterHttpMethod(s)
	}
}

func (s *HttpMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitHttpMethod(s)
	}
}

func (s *HttpMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitHttpMethod(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) HttpMethod() (localctx IHttpMethodContext) {
	localctx = NewHttpMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LoggingParserRULE_httpMethod)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LoggingParserPOST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.PostMethod()
		}


	case LoggingParserGET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.GetMethod()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IPostMethodContext is an interface to support dynamic dispatch.
type IPostMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostMethodContext differentiates from other interfaces.
	IsPostMethodContext()
}

type PostMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostMethodContext() *PostMethodContext {
	var p = new(PostMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_postMethod
	return p
}

func (*PostMethodContext) IsPostMethodContext() {}

func NewPostMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostMethodContext {
	var p = new(PostMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_postMethod

	return p
}

func (s *PostMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *PostMethodContext) POST() antlr.TerminalNode {
	return s.GetToken(LoggingParserPOST, 0)
}

func (s *PostMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PostMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterPostMethod(s)
	}
}

func (s *PostMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitPostMethod(s)
	}
}

func (s *PostMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitPostMethod(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) PostMethod() (localctx IPostMethodContext) {
	localctx = NewPostMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LoggingParserRULE_postMethod)

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
		p.SetState(74)
		p.Match(LoggingParserPOST)
	}



	return localctx
}


// IGetMethodContext is an interface to support dynamic dispatch.
type IGetMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetMethodContext differentiates from other interfaces.
	IsGetMethodContext()
}

type GetMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetMethodContext() *GetMethodContext {
	var p = new(GetMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_getMethod
	return p
}

func (*GetMethodContext) IsGetMethodContext() {}

func NewGetMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetMethodContext {
	var p = new(GetMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_getMethod

	return p
}

func (s *GetMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *GetMethodContext) GET() antlr.TerminalNode {
	return s.GetToken(LoggingParserGET, 0)
}

func (s *GetMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GetMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterGetMethod(s)
	}
}

func (s *GetMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitGetMethod(s)
	}
}

func (s *GetMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitGetMethod(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) GetMethod() (localctx IGetMethodContext) {
	localctx = NewGetMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LoggingParserRULE_getMethod)

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
		p.SetState(76)
		p.Match(LoggingParserGET)
	}



	return localctx
}


// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }
func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LoggingParserRULE_path)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(78)
			p.MatchWildcard()



		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}



	return localctx
}


// IPipeContext is an interface to support dynamic dispatch.
type IPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeContext differentiates from other interfaces.
	IsPipeContext()
}

type PipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeContext() *PipeContext {
	var p = new(PipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_pipe
	return p
}

func (*PipeContext) IsPipeContext() {}

func NewPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeContext {
	var p = new(PipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_pipe

	return p
}

func (s *PipeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(LoggingParserPIPE, 0)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (s *PipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitPipe(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) Pipe() (localctx IPipeContext) {
	localctx = NewPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LoggingParserRULE_pipe)

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
		p.Match(LoggingParserPIPE)
	}



	return localctx
}


// IStatusCodeContext is an interface to support dynamic dispatch.
type IStatusCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatusCodeContext differentiates from other interfaces.
	IsStatusCodeContext()
}

type StatusCodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatusCodeContext() *StatusCodeContext {
	var p = new(StatusCodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoggingParserRULE_statusCode
	return p
}

func (*StatusCodeContext) IsStatusCodeContext() {}

func NewStatusCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatusCodeContext {
	var p = new(StatusCodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoggingParserRULE_statusCode

	return p
}

func (s *StatusCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *StatusCodeContext) CODE() antlr.TerminalNode {
	return s.GetToken(LoggingParserCODE, 0)
}

func (s *StatusCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatusCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatusCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.EnterStatusCode(s)
	}
}

func (s *StatusCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoggingListener); ok {
		listenerT.ExitStatusCode(s)
	}
}

func (s *StatusCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LoggingVisitor:
		return t.VisitStatusCode(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LoggingParser) StatusCode() (localctx IStatusCodeContext) {
	localctx = NewStatusCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LoggingParserRULE_statusCode)

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
		p.SetState(86)
		p.Match(LoggingParserCODE)
	}



	return localctx
}


