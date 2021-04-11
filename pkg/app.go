package pkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Application struct {
	App  *fiber.App
	Conf Config
}

// NewApplication create a new application
func NewApplication(conf Config) Application {
	kbyte := 1024
	app := fiber.New(fiber.Config{
		Prefork:         conf.Server.Prefork,
		StrictRouting:   conf.Server.Strict,
		CaseSensitive:   conf.Server.Case,
		ETag:            conf.Server.Etag,
		BodyLimit:       conf.Server.BodyLimit * kbyte * kbyte,
		Concurrency:     conf.Server.Concurrency * kbyte,
		ReadTimeout:     conf.Server.Timeout.Read * time.Second,
		WriteTimeout:    conf.Server.Timeout.Write * time.Second,
		IdleTimeout:     conf.Server.Timeout.Idel * time.Second,
		ReadBufferSize:  conf.Server.Buffer.Read * kbyte,
		WriteBufferSize: conf.Server.Buffer.Write * kbyte,
	})
	return Application{
		App:  app,
		Conf: conf,
	}
}

// Run is start server
func (app *Application) Run() {
	app.Setup()

	app.App.Listen(fmt.Sprintf("%s:%d", app.Conf.App.Host, app.Conf.App.Port))
}

// Setup is setup init middleware
func (app *Application) Setup() {
	app.App.Use(logger.New())
	app.App.Use(recover.New())
	app.App.Use(cache.New(cache.Config{
		Expiration:   app.Conf.Cache.Expiration * time.Minute,
		CacheControl: app.Conf.Cache.Control,
	}))
}

// Environment check env
func (app *Application) Environment(env string) bool {
	return strings.ToLower(app.Conf.App.Env) == strings.ToLower(env)
}
