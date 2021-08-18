// Code generated from /Users/rafaeldurelli/Documents/development/golang/parserLogTest/Logging.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Logging

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by LoggingParser.
type LoggingVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LoggingParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by LoggingParser#log.
	VisitLog(ctx *LogContext) interface{}

	// Visit a parse tree produced by LoggingParser#gin.
	VisitGin(ctx *GinContext) interface{}

	// Visit a parse tree produced by LoggingParser#dateformat.
	VisitDateformat(ctx *DateformatContext) interface{}

	// Visit a parse tree produced by LoggingParser#requestTime.
	VisitRequestTime(ctx *RequestTimeContext) interface{}

	// Visit a parse tree produced by LoggingParser#httpMethod.
	VisitHttpMethod(ctx *HttpMethodContext) interface{}

	// Visit a parse tree produced by LoggingParser#postMethod.
	VisitPostMethod(ctx *PostMethodContext) interface{}

	// Visit a parse tree produced by LoggingParser#getMethod.
	VisitGetMethod(ctx *GetMethodContext) interface{}

	// Visit a parse tree produced by LoggingParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by LoggingParser#pipe.
	VisitPipe(ctx *PipeContext) interface{}

	// Visit a parse tree produced by LoggingParser#statusCode.
	VisitStatusCode(ctx *StatusCodeContext) interface{}

}