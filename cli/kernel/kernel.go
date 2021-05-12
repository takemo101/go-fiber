package kernel

import (
	"github.com/takemo101/go-fiber/cli/cmd"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// CLIBooter cli boot interface
type CLIBooter interface {
	CLIBoot()
}

// boot is initialize cli
func boot(
	lifecycle fx.Lifecycle,
	logger pkg.Logger,
	database pkg.Database,
	root cmd.RootCommand,
	booter CLIBooter,
) {
	sql, err := database.DB()
	if err != nil {
		logger.Info("database connection sql failed : %v", err)
	}

	defer sql.Close()

	logger.Info("-- start cli --")

	sql.SetMaxOpenConns(10)

	booter.CLIBoot()
	root.Cmd.Execute()

	logger.Info("-- stop cli --")
}

// app run
func Run(opts ...fx.Option) {
	newOpts := append(
		opts,
		pkg.Module,
		fx.Invoke(boot),
	)
	fx.New(newOpts...).Done()
}
