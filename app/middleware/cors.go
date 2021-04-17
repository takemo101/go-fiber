package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/takemo101/go-fiber/pkg"
)

// Cors is struct
type Cors struct {
	logger pkg.Logger
	app    pkg.Application
	config pkg.Config
}

// NewCors is create middleware
func NewCors(
	logger pkg.Logger,
	app pkg.Application,
	config pkg.Config,
) Cors {
	return Cors{
		logger: logger,
		app:    app,
		config: config,
	}
}

// Setup cors control middleware
func (m Cors) Setup() {
	m.app.App.Use(m.CreateHandler())
}

// CreateHandler is create middleware handler
func (m Cors) CreateHandler() fiber.Handler {
	m.logger.Info("setup cors")

	config := cors.Config{
		AllowMethods: strings.Join([]string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
			"HEAD",
		}, ", "),
		AllowHeaders: strings.Join([]string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Origin",
			"Cache-Control",
			"X-Requested-With",
		}, ", "),
		AllowCredentials: true,
		MaxAge:           int(m.config.Cors.MaxAge * time.Hour),
	}

	config.AllowOrigins = strings.Join(m.config.Cors.Origins, ", ")
	return cors.New(config)
}
