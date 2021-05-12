package main

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/cli/cmd"
	"github.com/takemo101/go-fiber/cli/kernel"
	"go.uber.org/fx"
)

func NewCLIBooter(
	commands cmd.Commands,
) kernel.CLIBooter {
	return commands
}

func main() {
	// boot cobra application
	kernel.Run(
		cmd.Module,
		app.Module,
		fx.Provide(NewCLIBooter),
	)
}
