package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// SessionAdminAuth admin auth
type SessionAdminAuth struct {
	logger pkg.Logger
	app    pkg.Application
	path   pkg.Path
	Auth   *support.SessionAdminAuth
}

// NewSessionAdminAuth is create middleware
func NewSessionAdminAuth(
	logger pkg.Logger,
	app pkg.Application,
	path pkg.Path,
	auth *support.SessionAdminAuth,
) SessionAdminAuth {
	return SessionAdminAuth{
		logger: logger,
		app:    app,
		path:   path,
		Auth:   auth,
	}
}

// Setup session admin auth middleware
func (m SessionAdminAuth) Setup() {
	m.logger.Info("setup session auth admin")
	m.app.App.Use(m.CreateHandler(true, "system/auth/login"))
}

// CreateHandler is create middleware handler
func (m SessionAdminAuth) CreateHandler(login bool, redirect string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := GetSession(c)
		if err != nil {
			return err
		}

		ok := m.Auth.AttemptSession(session)

		if (login && ok) || (!login && !ok) {
			return c.Next()
		}

		return c.Redirect(m.path.URL(redirect))
	}
}
