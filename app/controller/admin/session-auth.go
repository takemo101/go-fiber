package admin_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// SessionAuthController is index
type SessionAuthController struct {
	logger  pkg.Logger
	path    pkg.Path
	render  *helper.ViewRender
	auth    *support.SessionAdminAuth
	service service.AdminService
}

// NewSessionAuthController is create auth controller
func NewSessionAuthController(
	logger pkg.Logger,
	path pkg.Path,
	render *helper.ViewRender,
	auth *support.SessionAdminAuth,
	service service.AdminService,
) SessionAuthController {
	return SessionAuthController{
		logger:  logger,
		path:    path,
		render:  render,
		auth:    auth,
		service: service,
	}
}

// LoginForm render login form
func (ctl SessionAuthController) LoginForm(c *fiber.Ctx) error {
	return ctl.render.Render("auth/login", helper.DataMap{})
}

// Login login auth process
func (ctl SessionAuthController) Login(c *fiber.Ctx) error {
	var form form.Login
	redirect := "system/auth/login"

	session, sessionErr := middleware.GetSession(c)
	if sessionErr != nil {
		return ctl.render.Error(sessionErr)
	}

	if err := c.BodyParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	if err := form.Validate(func(email string, pass string) bool {
		return ctl.auth.Attempt(email, pass, session)
	}); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		return c.Redirect(ctl.path.URL(redirect))
	}

	return c.Redirect(ctl.path.URL("system"))
}

// Logout logout auth process
func (ctl SessionAuthController) Logout(c *fiber.Ctx) error {
	session, sessionErr := middleware.GetSession(c)
	if sessionErr != nil {
		return ctl.render.Error(sessionErr)
	}

	ctl.auth.Logout(session)
	return c.Redirect(ctl.path.URL("system/auth/login"))
}
