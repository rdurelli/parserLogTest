package routes

import (
	"github.com/rdurelli/loggingParser/src/controllers"
	"github.com/rdurelli/loggingParser/src/lib"
)

type LogRoutes struct {
	logger        lib.Logger
	handler       lib.RequestHandlerGin
	logController controllers.LogController
}

func (lR LogRoutes) Setup() {
	lR.logger.Info("Setting up LogRoutes")
	api := lR.handler.Gin.Group("/api")
	{
		api.GET("/logs", lR.logController.Find)
	}
}

func NewLogRoutes(logger lib.Logger, handler lib.RequestHandlerGin, logController controllers.LogController) LogRoutes {
	return LogRoutes{
		logger:        logger,
		handler:       handler,
		logController: logController,
	}
}
