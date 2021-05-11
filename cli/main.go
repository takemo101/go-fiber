package main

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/cli/cmd"
	"github.com/takemo101/go-fiber/cli/kernel"
)

func main() {
	// boot cobra application
	kernel.Run(
		cmd.Module,
		app.Module,
	)
}
