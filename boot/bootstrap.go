package boot

import (
	"context"

	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/route"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	pkg.Module,
	middleware.Module,
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
	middlewares middleware.Middlewares,
) {
	sql, err := database.DB()
	if err != nil {
		logger.Info("database connection sql failed : %v", err)
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("-- start application --")

			sql.SetMaxOpenConns(10)
			go func() {
				middlewares.Setup()
				routes.Setup()
				app.Run()
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("-- stop application --")
			sql.Close()
			return nil
		},
	})
}
