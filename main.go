package main

import (
	"github.com/takemo101/go-fiber/boot"
	"go.uber.org/fx"
)

func main() {
	// boot gin application
	fx.New(boot.Module).Run()
}
