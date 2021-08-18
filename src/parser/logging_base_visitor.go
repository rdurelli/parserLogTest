// Code generated from /Users/rafaeldurelli/Documents/development/golang/parserLogTest/Logging.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Logging

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLoggingVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLoggingVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitLog(ctx *LogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitGin(ctx *GinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitDateformat(ctx *DateformatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitRequestTime(ctx *RequestTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitHttpMethod(ctx *HttpMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitPostMethod(ctx *PostMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitGetMethod(ctx *GetMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitPipe(ctx *PipeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLoggingVisitor) VisitStatusCode(ctx *StatusCodeContext) interface{} {
	return v.VisitChildren(ctx)
}
