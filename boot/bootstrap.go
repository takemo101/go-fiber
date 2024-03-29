package boot

import (
	"context"

	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// AppBooter app boot interface
type AppBooter interface {
	AppBoot()
}

// boot is initialize application
func boot(
	lifecycle fx.Lifecycle,
	app pkg.Application,
	logger pkg.Logger,
	database pkg.Database,
	booter AppBooter,
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
				booter.AppBoot()
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

// app run
func Run(opts ...fx.Option) {
	newOpts := append(
		opts,
		pkg.Module,
		fx.Invoke(boot),
	)
	fx.New(newOpts...).Run()
}
