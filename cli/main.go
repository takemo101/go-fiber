package main

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/cli/cmd"
	"github.com/takemo101/go-fiber/cli/kernel"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

func NewCLIBooter(
	commands cmd.Commands,
) kernel.CLIBooter {
	return commands
}

func main() {

	// set config yml path
	pkg.ConfigPath = "config.yml"

	// boot cobra application
	kernel.Run(
		cmd.Module,
		app.Module,
		fx.Provide(NewCLIBooter),
	)
}
