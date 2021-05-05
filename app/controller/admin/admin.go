package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminController is index
type AdminController struct {
	logger  pkg.Logger
	path    pkg.Path
	render  *helper.ViewRender
	service service.AdminService
}

// NewAdminController is create admin controller
func NewAdminController(
	logger pkg.Logger,
	path pkg.Path,
	render *helper.ViewRender,
	service service.AdminService,
) AdminController {
	return AdminController{
		logger:  logger,
		path:    path,
		render:  render,
		service: service,
	}
}

// Index admin list
func (ctl AdminController) Index(c *fiber.Ctx) error {
	admins, err := ctl.service.Search()
	if err != nil {
		return err
	}
	return ctl.render.Render("admin/index", helper.DataMap{
		"admins": admins,
	})
}

// Create admin create form
func (ctl AdminController) Create(c *fiber.Ctx) error {
	return ctl.render.Render("admin/create", helper.DataMap{
		"content_footer": true,
	})
}

// Store admin store process
func (ctl AdminController) Store(c *fiber.Ctx) error {
	var form form.Admin

	if err := c.BodyParser(&form); err != nil {
		return err
	}

	if err := form.Validate(true); err != nil {
		middleware.SetSessionErrors(c, helper.ErrorsToMap(err))
		middleware.SetSessionInputs(c, helper.StructToFormMap(&form))
		return c.Redirect(ctl.path.URL("admin/admin/create"))
	}

	if _, err := ctl.service.Store(form); err != nil {
		return err
	}

	return c.Redirect(ctl.path.URL("admin/admin"))
}

// Edit admin edit form
func (ctl AdminController) Edit(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return convErr
	}

	admin, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return ctl.render.Error(findErr)
	}

	return ctl.render.Render("admin/edit", helper.DataMap{
		"admin":          admin,
		"content_footer": true,
	})
}

// Update admin update process
func (ctl AdminController) Update(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return convErr
	}

	var form form.Admin

	if err := c.BodyParser(&form); err != nil {
		return err
	}

	if err := form.Validate(false); err != nil {
		middleware.SetSessionErrors(c, helper.ErrorsToMap(err))
		middleware.SetSessionInputs(c, helper.StructToFormMap(&form))
		return c.Redirect(ctl.path.URL("admin/admin/" + c.Params("id") + "/edit"))
	}

	if _, err := ctl.service.Update(uint(id), form); err != nil {
		return err
	}

	return c.Redirect(ctl.path.URL("admin/admin/" + c.Params("id") + "/edit"))
}

// Delete admin delete process
func (ctl AdminController) Delete(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return convErr
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return err
	}

	return c.Redirect(ctl.path.URL("admin/admin"))
}
