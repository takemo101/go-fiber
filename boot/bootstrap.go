package boot

import (
	"context"

	"github.com/takemo101/go-fiber/app/route"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	pkg.Module,
	route.Module,
	fx.Invoke(boot),
)

// boot is initialize application
func boot(
	lifecycle fx.Lifecycle,
	app pkg.Application,
	logger pkg.Logger,
	database pkg.Database,
	routes route.Routes,
) {
	sql, err := database.DB.DB()
	if err != nil {
		logger.Zap.Info("database connection sql failed : %v", err)
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("-- start application --")

			sql.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				app.Run()
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("-- stop application --")
			sql.Close()
			return nil
		},
	})
}
