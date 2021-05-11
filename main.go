package main

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/boot"
	"go.uber.org/fx"
)

// NewBooter app create
func NewBooter(
	app app.AppModule,
) boot.Booter {
	return app
}

func main() {
	// boot gin application
	boot.Run(
		app.Module,
		fx.Provide(NewBooter),
	)
}
