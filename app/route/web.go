package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/pkg"
)

// WebRoute is struct
type WebRoute struct {
	logger pkg.Logger
	app    pkg.Application
}

// Setup is setup route
func (r WebRoute) Setup() {
	r.logger.Zap.Info("setup web-route")

	r.app.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

// NewWebRoute create new web route
func NewWebRoute(
	logger pkg.Logger,
	app pkg.Application,
) WebRoute {
	return WebRoute{
		logger: logger,
		app:    app,
	}
}
