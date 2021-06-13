package main

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/boot"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// NewAppBooter app create
func NewAppBooter(
	app app.AppModule,
) boot.AppBooter {
	return app
}

func main() {

	// set config yml path
	pkg.ConfigPath = "config.yml"

	// boot gin application
	boot.Run(
		app.Module,
		fx.Provide(NewAppBooter),
	)
}
