package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// JWTAuthController is jwt auth
type JWTAuthController struct {
	logger pkg.Logger
	config pkg.Config
	value  support.RequestValue
}

// NewJWTAuthController is create auth controller
func NewJWTAuthController(
	logger pkg.Logger,
	config pkg.Config,
	value support.RequestValue,
) JWTAuthController {
	return JWTAuthController{
		logger: logger,
		config: config,
		value:  value,
	}
}

// Login jwt auth
func (ctl JWTAuthController) Login(c *fiber.Ctx) error {
	var form form.Login
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.JsonError(c, err)
	}

	auth := ctl.value.GetJWTUserAuth(c)

	if err := form.Validate(func(email string, pass string) bool {
		return auth.Attempt(email, pass, c)
	}); err != nil {
		return response.JsonErrorMessages(c, err, helper.ErrorsToMap(err))
	}

	auth.GenerateTokenByUser(auth.User())

	token, err := auth.GenerateTokenByUser(auth.User())
	if err != nil {
		return response.JsonError(c, err)
	}

	return response.JsonSuccessWith(c, "login success", fiber.Map{
		"token": token,
	})
}

// TokenCheck token claims check
func (ctl JWTAuthController) TokenCheck(c *fiber.Ctx) error {
	auth := ctl.value.GetJWTUserAuth(c)
	response := ctl.value.GetResponseHelper(c)
	claims, err := auth.ExtractClaims(c)
	if err != nil {
		return response.JsonError(c, err)
	}

	return response.JsonSuccessWith(c, "success decode token", fiber.Map{
		"claims": claims,
	})
}

// LoginCheck login user check
func (ctl JWTAuthController) LoginCheck(c *fiber.Ctx) error {
	auth := ctl.value.GetJWTUserAuth(c)
	response := ctl.value.GetResponseHelper(c)
	return response.Json(c, fiber.Map{
		"id": auth.ID(),
	})
}
