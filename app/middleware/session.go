package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/takemo101/go-fiber/pkg"
)

// Session is struct
type Session struct {
	logger pkg.Logger
	app    pkg.Application
	config pkg.Config
}

// NewSession is create middleware
func NewSession(
	logger pkg.Logger,
	app pkg.Application,
	config pkg.Config,
) Session {
	return Session{
		logger: logger,
		app:    app,
		config: config,
	}
}

// Setup start session middleware
func (m Session) Setup() {
	m.logger.Info("setup session")
	store := session.New(session.Config{
		Expiration:     m.config.Session.Expiration,
		CookieName:     m.config.Session.Name,
		CookieDomain:   m.config.Session.Domain,
		CookiePath:     m.config.Session.Path,
		CookieSecure:   m.config.Session.Secure,
		CookieHTTPOnly: m.config.Session.HTTPOnly,
		KeyGenerator:   utils.UUID,
	})

	m.app.App.Use(func(c *fiber.Ctx) error {
		c.Locals("session", store)
		return c.Next()
	})
}
