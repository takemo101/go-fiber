package route

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	controller "github.com/takemo101/go-fiber/app/controller/api"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// ApiRoute is struct
type ApiRoute struct {
	logger         pkg.Logger
	app            pkg.Application
	cors           middleware.Cors
	jwt            middleware.JWTAuth
	authController controller.JWTAuthController
	value          support.RequestValue
}

// Setup is setup route
func (r ApiRoute) Setup() {
	r.logger.Info("setup api-route")

	app := r.app.App

	systemApi := app.Group("/system-api", r.cors.CreateHandler())
	{
		systemApi.Get("/", func(c *fiber.Ctx) error {
			response := r.value.GetResponseHelper(c)
			return response.Json(c, fiber.Map{
				"message": "it's system-api",
			})
		})
		systemApi.Get("/success", func(c *fiber.Ctx) error {
			response := r.value.GetResponseHelper(c)
			return response.JsonSuccessWith(c, "success", fiber.Map{
				"data": "json data",
			})
		})
		systemApi.Get("/error", func(c *fiber.Ctx) error {
			response := r.value.GetResponseHelper(c)
			return response.JsonErrorWith(c, errors.New("error"), fiber.Map{
				"data": "json data",
			})
		})
	}

	api := app.Group("/api", r.cors.CreateHandler())
	{
		// api test
		api.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "it's api",
			})
		})

		// login
		api.Post("/auth/login", r.authController.Login)
		// login check
		api.Get("/auth/check", r.authController.LoginCheck)

		// after auth
		auth := api.Group("/", r.jwt.CreateHandler())
		{
			// jwt token check
			auth.Get("/auth/check", r.authController.TokenCheck)
		}
	}
}

// NewApiRoute create new web route
func NewApiRoute(
	logger pkg.Logger,
	app pkg.Application,
	cors middleware.Cors,
	jwt middleware.JWTAuth,
	authController controller.JWTAuthController,
	value support.RequestValue,
) ApiRoute {
	return ApiRoute{
		logger:         logger,
		app:            app,
		cors:           cors,
		jwt:            jwt,
		authController: authController,
		value:          value,
	}
}
