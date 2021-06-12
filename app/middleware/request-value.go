package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// RequestValue is struct
type RequestValueInit struct {
	logger pkg.Logger
	app    pkg.Application
	value  support.RequestValue
	// value dependency
	adminRepository repository.AdminRepository
	config          pkg.Config
	path            pkg.Path
}

// NewRequestValueInit is create middleware
func NewRequestValueInit(
	logger pkg.Logger,
	app pkg.Application,
	// value dependency
	adminRepository repository.AdminRepository,
	config pkg.Config,
	path pkg.Path,
) RequestValueInit {
	return RequestValueInit{
		logger: logger,
		app:    app,
		// value dependency
		adminRepository: adminRepository,
		config:          config,
		path:            path,
	}
}

// Setup user-value control middleware
func (m RequestValueInit) Setup() {
	m.logger.Info("setup request-value init")
	m.app.App.Use(m.CreateHandler())
}

// CreateHandler is create middleware handler
func (m RequestValueInit) CreateHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// set SessionAdminAuth
		m.value.SetSessionAdminAuth(c, support.NewSessionAdminAuth(
			m.adminRepository,
		))
		// set ViewRender
		render := helper.NewViewRender()
		m.value.SetViewRender(c, render)
		// set ResponseHelper
		m.value.SetResponseHelper(c, helper.NewResponseHelper(
			m.logger,
			m.config,
			m.path,
			render,
		))
		return c.Next()
	}
}
