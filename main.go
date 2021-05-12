package main

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/boot"
	"go.uber.org/fx"
)

// NewAppBooter app create
func NewAppBooter(
	app app.AppModule,
) boot.AppBooter {
	return app
}

func main() {
	// boot gin application
	boot.Run(
		app.Module,
		fx.Provide(NewAppBooter),
	)
}
