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
	logs, err := lC.service.Find(offset, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if len(*logs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "There is not log to get in the database. Please, call support.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"logs":    *logs,
		"size":    len(*logs),
		"success": true,
	})
}
