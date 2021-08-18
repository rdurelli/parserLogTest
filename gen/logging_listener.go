// Code generated from /Users/rafaeldurelli/Documents/development/golang/parserLogTest/Logging.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Logging

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LoggingListener is a complete listener for a parse tree produced by LoggingParser.
type LoggingListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterLog is called when entering the log production.
	EnterLog(c *LogContext)

	// EnterGin is called when entering the gin production.
	EnterGin(c *GinContext)

	// EnterDateformat is called when entering the dateformat production.
	EnterDateformat(c *DateformatContext)

	// EnterRequestTime is called when entering the requestTime production.
	EnterRequestTime(c *RequestTimeContext)

	// EnterHttpMethod is called when entering the httpMethod production.
	EnterHttpMethod(c *HttpMethodContext)

	// EnterPostMethod is called when entering the postMethod production.
	EnterPostMethod(c *PostMethodContext)

	// EnterGetMethod is called when entering the getMethod production.
	EnterGetMethod(c *GetMethodContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

	// EnterStatusCode is called when entering the statusCode production.
	EnterStatusCode(c *StatusCodeContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitLog is called when exiting the log production.
	ExitLog(c *LogContext)

	// ExitGin is called when exiting the gin production.
	ExitGin(c *GinContext)

	// ExitDateformat is called when exiting the dateformat production.
	ExitDateformat(c *DateformatContext)

	// ExitRequestTime is called when exiting the requestTime production.
	ExitRequestTime(c *RequestTimeContext)

	// ExitHttpMethod is called when exiting the httpMethod production.
	ExitHttpMethod(c *HttpMethodContext)

	// ExitPostMethod is called when exiting the postMethod production.
	ExitPostMethod(c *PostMethodContext)

	// ExitGetMethod is called when exiting the getMethod production.
	ExitGetMethod(c *GetMethodContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitStatusCode is called when exiting the statusCode production.
	ExitStatusCode(c *StatusCodeContext)
}
