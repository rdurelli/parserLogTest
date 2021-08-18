package bootstrap

import (
	"context"
	"time"

	"github.com/rdurelli/loggingParser/src/routes"

	"github.com/rdurelli/loggingParser/src/controllers"
	"github.com/rdurelli/loggingParser/src/services"

	"github.com/rdurelli/loggingParser/src/models"

	"github.com/rdurelli/loggingParser/src/utils"

	"github.com/rdurelli/loggingParser/src/lib"
	"github.com/rdurelli/loggingParser/src/repositories"
	"go.uber.org/fx"
)

var Module = fx.Options(
	lib.Module,
	repositories.Module,
	controllers.Module,
	routes.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle,
	env lib.Env,
	routes routes.Routes,
	logger lib.Logger,
	db lib.Db,
	handler lib.RequestHandlerGin) {

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting Application")
			logger.Info("---------------------")
			logger.Info("------- CLEAN -------")
			logger.Info("---------------------")
			db.Db.SetMaxOpenConns(10)
			logs := make([]models.Log, 0)
			readerLog := utils.NewReaderLog(env.LogPath, &logs, db, logger)
			go func() {
				s := time.Now()
				readerLog.Execute()
				logger.Info("Time taken to process all logs ", time.Since(s))
				//fmt.Println("\nTime taken - ", time.Since(s))
			}()
			go func() {
				//middlewares.Setup()
				routes.Setup()
				handler.Gin.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping Application")
			db.Db.Close()
			return nil
		},
	})

}
