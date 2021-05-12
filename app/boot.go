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
	fx.Provide(NewAppModule),
)

// AppModule is module root struct
type AppModule struct {
	routes      route.Routes
	middlewares middleware.Middlewares
}

func NewAppModule(
	routes route.Routes,
	middlewares middleware.Middlewares,
) AppModule {
	return AppModule{
		routes:      routes,
		middlewares: middlewares,
	}
}

// AppBoot all setup
func (module AppModule) AppBoot() {
	module.middlewares.Setup()
	module.routes.Setup()
}
