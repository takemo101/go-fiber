package kernel

import (
	"github.com/takemo101/go-fiber/cli/cmd"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	pkg.Module,
	cmd.Module,
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
	sql, err := database.DB.DB()
	if err != nil {
		logger.Zap.Info("database connection sql failed : %v", err)
	}

	defer sql.Close()

	logger.Zap.Info("-- start cli --")

	sql.SetMaxOpenConns(10)

	commands.Setup()
	root.Cmd.Execute()

	logger.Zap.Info("-- stop cli --")
}
