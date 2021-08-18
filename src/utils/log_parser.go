package utils

import (
	"github.com/rdurelli/loggingParser/src/models"
	"github.com/rdurelli/loggingParser/src/repositories"

	"github.com/rdurelli/loggingParser/src/parser"
)

type LogParser struct {
	*parser.BaseLoggingListener
	Logs       *[]models.Log
	Repository repositories.Repository
}

var logToPersist = models.Log{}

func (s *LogParser) EnterStart(ctx *parser.StartContext) {
	logToPersist = models.Log{}
}

// EnterPostMethod is called when production postMethod is entered.
func (s *LogParser) EnterPostMethod(ctx *parser.PostMethodContext) {
	logToPersist.HttpMethod = ctx.GetText()
}

// EnterGetMethod is called when production getMethod is entered.
func (s *LogParser) EnterGetMethod(ctx *parser.GetMethodContext) {
	logToPersist.HttpMethod = ctx.GetText()
}

// EnterStatusCode is called when production statusCode is entered.
func (s *LogParser) EnterStatusCode(ctx *parser.StatusCodeContext) {
	logToPersist.HttpCode = ctx.GetText()
}

// EnterDateformat is called when production dateformat is entered.
func (s *LogParser) EnterDateformat(ctx *parser.DateformatContext) {
	logToPersist.Date = ctx.GetText()
}

// EnterPath is called when production path is entered.
func (s *LogParser) EnterPath(ctx *parser.PathContext) {
	logToPersist.Path = ctx.GetText()
}

// ExitPath is called when production path is exited.
func (s *LogParser) ExitPath(ctx *parser.PathContext) {
	if logToPersist.HttpMethod == "POST" || logToPersist.HttpMethod == "post" || logToPersist.HttpMethod == "GET" || logToPersist.HttpMethod == "get" {
		*s.Logs = append(*s.Logs, logToPersist)
		s.Repository.Save(&logToPersist)
	}
	//if len(*s.Logs) == 100 {
	//	s.Repository.SaveBatch(s.Logs)
	//}
}
