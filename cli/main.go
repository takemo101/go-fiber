package main

import (
	"github.com/takemo101/go-fiber/cli/kernel"
	"go.uber.org/fx"
)

func main() {
	// boot cobra application
	fx.New(kernel.Module).Done()
}
