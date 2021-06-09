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

// @title GoFiber
// @version 1.0
// @description GoFiber Api Doc
// @termsOfService http://swagger.io/terms/
// @contact.name takemo
// @contact.email takemo101@gmail.com
// @host localhost:8000
// @BasePath /api
func main() {
	// boot gin application
	boot.Run(
		app.Module,
		fx.Provide(NewAppBooter),
	)
}
