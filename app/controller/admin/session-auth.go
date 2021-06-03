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
	logger  pkg.Logger
	value   support.RequestValue
	service service.AdminService
}

// NewSessionAuthController is create auth controller
func NewSessionAuthController(
	logger pkg.Logger,
	value support.RequestValue,
	service service.AdminService,
) SessionAuthController {
	return SessionAuthController{
		logger:  logger,
		value:   value,
		service: service,
	}
}

// LoginForm render login form
func (ctl SessionAuthController) LoginForm(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	return response.View("auth/login", helper.DataMap{})
}

// Login login auth process
func (ctl SessionAuthController) Login(c *fiber.Ctx) error {
	var form form.Login
	response := ctl.value.GetResponseHelper(c)

	session, sessionErr := middleware.GetSession(c)
	if sessionErr != nil {
		return response.Error(sessionErr)
	}

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	auth := ctl.value.GetSessionAdminAuth(c)

	if err := form.Validate(func(email string, pass string) bool {
		return auth.Attempt(email, pass, session)
	}); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		return response.Back(c)
	}

	return response.Redirect(c, "system")
}

// Logout logout auth process
func (ctl SessionAuthController) Logout(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	session, sessionErr := middleware.GetSession(c)
	if sessionErr != nil {
		return response.Error(sessionErr)
	}

	auth := ctl.value.GetSessionAdminAuth(c)

	auth.Logout(session)
	return response.Redirect(c, "system/auth/login")
}
