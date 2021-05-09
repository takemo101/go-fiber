package admin_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// UserController is index
type UserController struct {
	logger  pkg.Logger
	path    pkg.Path
	render  *helper.ViewRender
	service service.UserService
}

// NewUserController is create user controller
func NewUserController(
	logger pkg.Logger,
	path pkg.Path,
	render *helper.ViewRender,
	service service.UserService,
) UserController {
	return UserController{
		logger:  logger,
		path:    path,
		render:  render,
		service: service,
	}
}

// Index render user list
func (ctl UserController) Index(c *fiber.Ctx) error {
	var form form.UserSearch

	if err := c.QueryParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	users, err := ctl.service.Search(form, 20)
	if err != nil {
		return ctl.render.Error(err)
	}
	return ctl.render.Render("user/index", helper.DataMap{
		"users": users,
	})
}

// Create render user create form
func (ctl UserController) Create(c *fiber.Ctx) error {
	return ctl.render.Render("user/create", helper.DataMap{
		"content_footer": true,
	})
}

// Store user store process
func (ctl UserController) Store(c *fiber.Ctx) error {
	var form form.User

	if err := c.BodyParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	if err := form.Validate(true, 0, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return c.Redirect(ctl.path.URL("system/user/create"))
	}

	if _, err := ctl.service.Store(form); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message())
	return c.Redirect(ctl.path.URL("system/user"))
}

// Edit render user edit form
func (ctl UserController) Edit(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.render.Error(convErr)
	}

	user, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return ctl.render.Error(findErr)
	}

	return ctl.render.Render("user/edit", helper.DataMap{
		"user":           user,
		"content_footer": true,
	})
}

// Update user update process
func (ctl UserController) Update(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.render.Error(convErr)
	}

	var form form.User

	if err := c.BodyParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	if err := form.Validate(false, uint(id), ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return c.Redirect(ctl.path.URL("system/user/%s/edit", c.Params("id")))
	}

	if _, err := ctl.service.Update(uint(id), form); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message())
	return c.Redirect(ctl.path.URL("system/user/%s/edit", c.Params("id")))
}

// Delete user delete process
func (ctl UserController) Delete(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.render.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message())
	return c.Redirect(ctl.path.URL("system/user"))
}
