package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminRoute is struct
type AdminRoute struct {
	logger pkg.Logger
	app    pkg.Application
	csrf   middleware.Csrf
}

// Setup is setup route
func (r AdminRoute) Setup() {
	r.logger.Info("setup admin-route")

	app := r.app.App

	admin := app.Group(
		"/admin",
		r.csrf.CreateHandler(),
	)
	{
		admin.Get("/", func(c *fiber.Ctx) error {
			return c.Render("home", fiber.Map{
				"Title": "Dashboard",
			})
		})
		admin.Get("/user", func(c *fiber.Ctx) error {
			return c.Render("user/index", fiber.Map{
				"Title": "User",
			})
		})
	}
}

// NewAdminRoute create new admin route
func NewAdminRoute(
	logger pkg.Logger,
	app pkg.Application,
	csrf middleware.Csrf,
) AdminRoute {
	return AdminRoute{
		logger: logger,
		app:    app,
		csrf:   csrf,
	}
}
