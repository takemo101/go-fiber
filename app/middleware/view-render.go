package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// ViewRender is struct
type ViewRender struct {
	logger pkg.Logger
	value  support.RequestValue
}

// NewViewRender is create middleware
func NewViewRender(
	logger pkg.Logger,
	value support.RequestValue,
) ViewRender {
	return ViewRender{
		logger: logger,
		value:  value,
	}
}

func (v ViewRender) CreateHandler(handler func(*fiber.Ctx, *helper.ViewRender)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		render := v.value.GetViewRender(c)
		return render.HandleRender(c, handler)
	}
}
