package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// JWTAuth is struct
type JWTAuth struct {
	logger pkg.Logger
	config pkg.Config
	app    pkg.Application
	value  support.RequestValue
}

// NewJWTAuth is create middleware
func NewJWTAuth(
	logger pkg.Logger,
	config pkg.Config,
	app pkg.Application,
	value support.RequestValue,
) JWTAuth {
	return JWTAuth{
		logger: logger,
		config: config,
		app:    app,
		value:  value,
	}
}

// Setup jwt control middleware
func (m JWTAuth) Setup() {
	m.app.App.Use(m.CreateHandler())
}

// CreateHandler is create middleware handler
func (m JWTAuth) CreateHandler() fiber.Handler {
	m.logger.Info("setup jwt auth")

	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(m.config.JWT.Signing.Key),
		SigningMethod:  strings.ToUpper(m.config.JWT.Signing.Method),
		ContextKey:     m.config.JWT.Context.Key,
		TokenLookup:    m.config.JWT.Lookup,
		AuthScheme:     m.config.JWT.Scheme,
		ErrorHandler:   m.ErrorHandler,
		SuccessHandler: m.SuccessHandler,
	})
}

// ErrorHandler
func (m JWTAuth) ErrorHandler(c *fiber.Ctx, err error) error {
	response := m.value.GetResponseHelper(c)
	if err.Error() == "Missing or malformed JWT" {
		return response.JsonErrorSimple(c, err, fiber.StatusBadRequest)
	}
	c.Status(fiber.StatusUnauthorized)
	return response.JsonErrorSimple(c, err, fiber.StatusUnauthorized)
}

// SuccessHandler
func (m JWTAuth) SuccessHandler(c *fiber.Ctx) error {
	auth := m.value.GetJWTUserAuth(c)
	response := m.value.GetResponseHelper(c)
	if auth.AttemptToken(c) {
		return c.Next()
	}
	return response.JsonErrorSimple(c, errors.New("not found token user"), fiber.StatusBadRequest)
}
