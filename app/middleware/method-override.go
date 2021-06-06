package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/pkg"
)

// MethodOverride is struct
type MethodOverride struct {
	logger pkg.Logger
	app    pkg.Application
}

// NewMethodOverride is create middleware
func NewMethodOverride(
	logger pkg.Logger,
	app pkg.Application,
) MethodOverride {
	return MethodOverride{
		logger: logger,
		app:    app,
	}
}

// Setup request method override control middleware
func (m MethodOverride) Setup() {
	m.logger.Info("setup request-method override")

	m.app.App.Use(func(c *fiber.Ctx) error {
		method := strings.ToUpper(c.FormValue("_method"))
		if c.Method() != fiber.MethodPost {
			return c.Next()
		}

		switch method {
		case fiber.MethodDelete:
			c.Method(fiber.MethodDelete)
		case fiber.MethodPut:
			c.Method(fiber.MethodPut)
		case fiber.MethodPatch:
			c.Method(fiber.MethodPatch)
		}

		return c.Next()
	})
}
