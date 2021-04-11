package route

import (
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Module export
var Module = fx.Options(
	fx.Provide(NewWebRoute),
	fx.Provide(NewRoute),
)

// Routes is slice
type Routes []Route

// Route is interface
type Route interface {
	Setup()
}

// NewRoute is setup routes
func NewRoute(
	config pkg.Config,
	web WebRoute,
) Routes {
	return Routes{
		web,
	}
}

// Setup all the route
func (routes Routes) Setup() {
	for _, route := range routes {
		route.Setup()
	}
}
