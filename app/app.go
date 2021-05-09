package app

import (
	"github.com/takemo101/go-fiber/app/controller"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/app/route"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"go.uber.org/fx"
)

var Module = fx.Options(
	middleware.Module,
	helper.Module,
	repository.Module,
	query.Module,
	service.Module,
	support.Module,
	route.Module,
	controller.Module,
	fx.Provide(NewMainModule),
)

// MainAppModule is module root struct
type MainModule struct {
	routes      route.Routes
	middlewares middleware.Middlewares
}

func NewMainModule(
	routes route.Routes,
	middlewares middleware.Middlewares,
) MainModule {
	return MainModule{
		routes:      routes,
		middlewares: middlewares,
	}
}

// Boot all setup
func (main MainModule) Boot() {
	main.middlewares.Setup()
	main.routes.Setup()
}
