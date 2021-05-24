package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// SessionAuthController is session auth
type SessionAuthController struct {
	logger   pkg.Logger
	response *helper.ResponseHelper
	auth     *support.SessionAdminAuth
	service  service.AdminService
}

// NewSessionAuthController is create auth controller
func NewSessionAuthController(
	logger pkg.Logger,
	response *helper.ResponseHelper,
	auth *support.SessionAdminAuth,
	service service.AdminService,
) SessionAuthController {
	return SessionAuthController{
		logger:   logger,
		response: response,
		auth:     auth,
		service:  service,
	}
}

// LoginForm render login form
func (ctl SessionAuthController) LoginForm(c *fiber.Ctx) error {
	return ctl.response.View("auth/login", helper.DataMap{})
}

// Login login auth process
func (ctl SessionAuthController) Login(c *fiber.Ctx) error {
	var form form.Login

	session, sessionErr := middleware.GetSession(c)
	if sessionErr != nil {
		return ctl.response.Error(sessionErr)
	}

	if err := c.BodyParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	if err := form.Validate(func(email string, pass string) bool {
		return ctl.auth.Attempt(email, pass, session)
	}); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		return ctl.response.Back(c)
	}

	return ctl.response.Redirect(c, "system")
}

// Logout logout auth process
func (ctl SessionAuthController) Logout(c *fiber.Ctx) error {
	session, sessionErr := middleware.GetSession(c)
	if sessionErr != nil {
		return ctl.response.Error(sessionErr)
	}

	ctl.auth.Logout(session)
	return ctl.response.Redirect(c, "system/auth/login")
}
