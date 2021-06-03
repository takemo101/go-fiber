package route

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// ApiRoute is struct
type ApiRoute struct {
	logger pkg.Logger
	app    pkg.Application
	cors   middleware.Cors
	value  support.RequestValue
}

// Setup is setup route
func (r ApiRoute) Setup() {
	r.logger.Info("setup web-route")

	app := r.app.App

	api := app.Group("/api", r.cors.CreateHandler())
	{
		api.Get("/", func(c *fiber.Ctx) error {
			response := r.value.GetResponseHelper(c)
			return response.Json(c, fiber.Map{
				"message": "it's api",
			})
		})
		api.Get("/success", func(c *fiber.Ctx) error {
			response := r.value.GetResponseHelper(c)
			return response.JsonSuccessWith(c, "success", fiber.Map{
				"data": "json data",
			})
		})
		api.Get("/error", func(c *fiber.Ctx) error {
			response := r.value.GetResponseHelper(c)
			return response.JsonErrorWith(c, errors.New("error"), fiber.Map{
				"data": "json data",
			})
		})
	}
}

// NewApiRoute create new web route
func NewApiRoute(
	logger pkg.Logger,
	app pkg.Application,
	cors middleware.Cors,
	value support.RequestValue,
) ApiRoute {
	return ApiRoute{
		logger: logger,
		app:    app,
		cors:   cors,
		value:  value,
	}
}
