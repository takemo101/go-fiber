package boot

import (
	"context"

	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/module"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	pkg.Module,
	app.Module,
	module.Module,
	fx.Invoke(boot),
)

// boot is initialize application
func boot(
	lifecycle fx.Lifecycle,
	app pkg.Application,
	logger pkg.Logger,
	database pkg.Database,
	main app.MainModule,
	modules module.ModuleParts,
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
				main.Boot()
				modules.Boot()
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
