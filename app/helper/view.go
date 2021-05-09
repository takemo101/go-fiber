package helper

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"github.com/takemo101/go-fiber/pkg"
)

type DataMap map[string]interface{}

// ViewRender is render manage
type ViewRender struct {
	logger pkg.Logger
	data   *ViewData
}

// ViewData is render data
type ViewData struct {
	data DataMap
	js   DataMap
	name string
	err  error
}

func NewViewRender(
	logger pkg.Logger,
) *ViewRender {
	return &ViewRender{
		logger: logger,
	}
}

func (v *ViewRender) Render(name string, data DataMap) error {
	v.SetName(name)
	v.SetData(data)
	return nil
}

func (v *ViewRender) SetData(data DataMap) {
	mergo.Merge(&v.data.data, data)
}

func (v *ViewRender) SetJS(data DataMap) {
	mergo.Merge(&v.data.js, data)
}

func (v *ViewRender) SetName(name string) {
	v.data.name = name
}

func (v *ViewRender) Error(err error) error {
	v.data.err = err
	v.logger.Error(err)
	v.SetName("error/error")
	return nil
}

func (v *ViewRender) Next(c *fiber.Ctx, handler func(*fiber.Ctx, *ViewRender)) error {
	v.data = new(ViewData)
	err := c.Next()
	if err == nil && len(v.data.name) > 0 {
		handler(c, v)

		js, err := json.Marshal(v.data.js)
		if err != nil {
			return err
		}

		v.data.data["js"] = string(js)

		if v.data.err == nil {
			return c.Render(v.data.name, fiber.Map(v.data.data))
		} else {
			v.data.data["error_message"] = v.data.err.Error()
			c.Status(400)
			return c.Render(v.data.name, fiber.Map(v.data.data))
		}
	}
	return err
}

func (v *ViewRender) CreateHandler(handler func(*fiber.Ctx, *ViewRender)) fiber.Handler {
	v.logger.Info("setup view-render")

	return func(c *fiber.Ctx) error {
		return v.Next(c, handler)
	}
}
