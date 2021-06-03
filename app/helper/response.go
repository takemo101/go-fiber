package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"github.com/takemo101/go-fiber/pkg"
)

// ResponseHelper response helper
type ResponseHelper struct {
	path   pkg.Path
	render *ViewRender
}

// NewResponseHelper response utility
func NewResponseHelper(
	path pkg.Path,
	render *ViewRender,
) *ResponseHelper {
	return &ResponseHelper{
		path:   path,
		render: render,
	}
}

// Back response redirect back
func (helper *ResponseHelper) Back(c *fiber.Ctx) error {
	back := string(c.Request().Header.Referer())
	return c.Redirect(back)
}

// Redirect response redirect
func (helper *ResponseHelper) Redirect(c *fiber.Ctx, path string) error {
	return c.Redirect(helper.path.URL(path))
}

// Json response json
func (helper *ResponseHelper) Json(c *fiber.Ctx, data fiber.Map) error {
	return c.JSON(data)
}

// JsonSuccess response json success
func (helper *ResponseHelper) JsonSuccess(c *fiber.Ctx, message string) error {
	return helper.Json(c, fiber.Map{
		"success": true,
		"message": message,
	})
}

// JsonSuccessWith response json success with data
func (helper *ResponseHelper) JsonSuccessWith(c *fiber.Ctx, message string, data fiber.Map) error {
	mainData := fiber.Map{
		"success": true,
		"message": message,
	}
	mergo.Merge(
		&mainData,
		data,
	)
	return helper.Json(c, mainData)
}

// JsonErrorSimple response json error
func (helper *ResponseHelper) JsonErrorSimple(c *fiber.Ctx, err error) error {
	return helper.Json(c, fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}

// JsonError response json error
func (helper *ResponseHelper) JsonError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusInternalServerError)
	return helper.JsonErrorSimple(c, err)
}

// JsonErrorWith response json error with data
func (helper *ResponseHelper) JsonErrorWith(c *fiber.Ctx, err error, data fiber.Map) error {
	mainData := fiber.Map{
		"success": false,
		"error":   err.Error(),
	}
	mergo.Merge(
		&mainData,
		data,
	)
	c.Status(fiber.StatusInternalServerError)
	return helper.Json(c, mainData)
}

// JsonErrorMessages response json error with error_messages
func (helper *ResponseHelper) JsonErrorMessages(c *fiber.Ctx, err error, messages map[string]string) error {
	c.Status(fiber.StatusUnprocessableEntity)
	return helper.Json(c, fiber.Map{
		"success":        false,
		"error":          err.Error(),
		"error_messages": messages,
	})
}

// JS set template javascript data
func (helper *ResponseHelper) JS(data DataMap) {
	helper.render.SetJS(data)
}

// Data set template view data
func (helper *ResponseHelper) Data(data DataMap) {
	helper.render.SetData(data)
}

// View render template view
func (helper *ResponseHelper) View(name string, data DataMap) error {
	return helper.render.Render(name, data)
}

// Error render template error
func (helper *ResponseHelper) Error(err error) error {
	return helper.render.Error(err)
}
