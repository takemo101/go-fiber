package kernel

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/cli/cmd"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	pkg.Module,
	cmd.Module,
	app.Module,
	fx.Invoke(boot),
)

// boot is initialize cli
func boot(
	lifecycle fx.Lifecycle,
	logger pkg.Logger,
	database pkg.Database,
	commands cmd.Commands,
	root cmd.RootCommand,
) {
	sql, err := database.DB()
	if err != nil {
		logger.Info("database connection sql failed : %v", err)
	}

	defer sql.Close()

	logger.Info("-- start cli --")

	sql.SetMaxOpenConns(10)

	commands.Setup()
	root.Cmd.Execute()

	logger.Info("-- stop cli --")
}
