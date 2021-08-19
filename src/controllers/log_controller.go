package controllers

import (
	"net/http"

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
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	if offset == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please, provide the offset",
		})
		return
	}
	logs, err := lC.service.Find(offset, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"logs": *logs,
		"size": len(*logs),
	})
}
