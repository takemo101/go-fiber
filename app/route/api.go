package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/pkg"
)

// ApiRoute is struct
type ApiRoute struct {
	logger pkg.Logger
	app    pkg.Application
	cors   middleware.Cors
}

// Setup is setup route
func (r ApiRoute) Setup() {
	r.logger.Info("setup web-route")

	app := r.app.App

	api := app.Group("/api/v1", r.cors.CreateHandler())
	{
		api.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "its api",
			})
		})
	}
}

// NewApiRoute create new web route
func NewApiRoute(
	logger pkg.Logger,
	app pkg.Application,
	cors middleware.Cors,
) ApiRoute {
	return ApiRoute{
		logger: logger,
		app:    app,
		cors:   cors,
	}
}
