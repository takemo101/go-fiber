package route

import (
	"github.com/takemo101/go-fiber/pkg"
	"github.com/takemo101/go-fiber/pkg/contract"
	"go.uber.org/fx"
)

// Module export
var Module = fx.Options(
	fx.Provide(NewApiRoute),
	fx.Provide(NewAdminRoute),
	fx.Provide(NewRoute),
)

// Routes is slice
type Routes []contract.Route

// NewRoute is setup routes
func NewRoute(
	config pkg.Config,
	api ApiRoute,
	admin AdminRoute,
) Routes {
	return Routes{
		api,
		admin,
	}
}

// Setup all the route
func (routes Routes) Setup() {
	for _, route := range routes {
		route.Setup()
	}
}
