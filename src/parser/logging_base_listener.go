// Code generated from /Users/rafaeldurelli/Documents/development/golang/parserLogTest/Logging.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Logging

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLoggingListener is a complete listener for a parse tree produced by LoggingParser.
type BaseLoggingListener struct{}

var _ LoggingListener = &BaseLoggingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLoggingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLoggingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLoggingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLoggingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseLoggingListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseLoggingListener) ExitStart(ctx *StartContext) {}

// EnterLog is called when production log is entered.
func (s *BaseLoggingListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseLoggingListener) ExitLog(ctx *LogContext) {}

// EnterGin is called when production gin is entered.
func (s *BaseLoggingListener) EnterGin(ctx *GinContext) {}

// ExitGin is called when production gin is exited.
func (s *BaseLoggingListener) ExitGin(ctx *GinContext) {}

// EnterDateformat is called when production dateformat is entered.
func (s *BaseLoggingListener) EnterDateformat(ctx *DateformatContext) {}

// ExitDateformat is called when production dateformat is exited.
func (s *BaseLoggingListener) ExitDateformat(ctx *DateformatContext) {}

// EnterRequestTime is called when production requestTime is entered.
func (s *BaseLoggingListener) EnterRequestTime(ctx *RequestTimeContext) {}

// ExitRequestTime is called when production requestTime is exited.
func (s *BaseLoggingListener) ExitRequestTime(ctx *RequestTimeContext) {}

// EnterHttpMethod is called when production httpMethod is entered.
func (s *BaseLoggingListener) EnterHttpMethod(ctx *HttpMethodContext) {}

// ExitHttpMethod is called when production httpMethod is exited.
func (s *BaseLoggingListener) ExitHttpMethod(ctx *HttpMethodContext) {}

// EnterPostMethod is called when production postMethod is entered.
func (s *BaseLoggingListener) EnterPostMethod(ctx *PostMethodContext) {}

// ExitPostMethod is called when production postMethod is exited.
func (s *BaseLoggingListener) ExitPostMethod(ctx *PostMethodContext) {}

// EnterGetMethod is called when production getMethod is entered.
func (s *BaseLoggingListener) EnterGetMethod(ctx *GetMethodContext) {}

// ExitGetMethod is called when production getMethod is exited.
func (s *BaseLoggingListener) ExitGetMethod(ctx *GetMethodContext) {}

// EnterPath is called when production path is entered.
func (s *BaseLoggingListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseLoggingListener) ExitPath(ctx *PathContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BaseLoggingListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BaseLoggingListener) ExitPipe(ctx *PipeContext) {}

// EnterStatusCode is called when production statusCode is entered.
func (s *BaseLoggingListener) EnterStatusCode(ctx *StatusCodeContext) {}

// ExitStatusCode is called when production statusCode is exited.
func (s *BaseLoggingListener) ExitStatusCode(ctx *StatusCodeContext) {}
