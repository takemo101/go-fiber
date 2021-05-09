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
	logger   pkg.Logger
	response *helper.ResponseHelper
	service  service.UserService
}

// NewUserController is create user controller
func NewUserController(
	logger pkg.Logger,
	response *helper.ResponseHelper,
	service service.UserService,
) UserController {
	return UserController{
		logger:   logger,
		response: response,
		service:  service,
	}
}

// Index render user list
func (ctl UserController) Index(c *fiber.Ctx) error {
	var form form.UserSearch

	if err := c.QueryParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	users, err := ctl.service.Search(form, 20)
	if err != nil {
		return ctl.response.Error(err)
	}
	return ctl.response.View("user/index", helper.DataMap{
		"users": users,
	})
}

// Create render user create form
func (ctl UserController) Create(c *fiber.Ctx) error {
	return ctl.response.View("user/create", helper.DataMap{
		"content_footer": true,
	})
}

// Store user store process
func (ctl UserController) Store(c *fiber.Ctx) error {
	var form form.User

	if err := c.BodyParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	if err := form.Validate(true, 0, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return ctl.response.Back(c)
	}

	if _, err := ctl.service.Store(form); err != nil {
		return ctl.response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message())
	return ctl.response.Redirect(c, "system/user")
}

// Edit render user edit form
func (ctl UserController) Edit(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.response.Error(convErr)
	}

	user, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return ctl.response.Error(findErr)
	}

	return ctl.response.View("user/edit", helper.DataMap{
		"user":           user,
		"content_footer": true,
	})
}

// Update user update process
func (ctl UserController) Update(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.response.Error(convErr)
	}

	var form form.User

	if err := c.BodyParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	if err := form.Validate(false, uint(id), ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return ctl.response.Back(c)
	}

	if _, err := ctl.service.Update(uint(id), form); err != nil {
		return ctl.response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message())
	return ctl.response.Back(c)
}

// Delete user delete process
func (ctl UserController) Delete(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return ctl.response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message())
	return ctl.response.Redirect(c, "system/user")
}
