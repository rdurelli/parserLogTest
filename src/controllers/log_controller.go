package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rdurelli/loggingParser/src/lib"
	"github.com/rdurelli/loggingParser/src/services"
)

type LogController struct {
	logger  lib.Logger
	service services.LogService
}

func NewLogController(logger lib.Logger, service services.LogService) LogController {
	return LogController{
		logger:  logger,
		service: service,
	}
}

func (lC LogController) Find(c *gin.Context) {
	panic("implement me")
}
