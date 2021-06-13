package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
)

// UserController is user
type UserController struct {
	value   support.RequestValue
	service service.UserService
}

// NewUserController is create user controller
func NewUserController(
	value support.RequestValue,
	service service.UserService,
) UserController {
	return UserController{
		value:   value,
		service: service,
	}
}

// Index render user list
func (ctl UserController) Index(c *fiber.Ctx) error {
	var form form.UserSearch
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	users, paginator, err := ctl.service.Search(object.NewUserSearchInput(
		form.Keyword,
		form.Page,
	), 20)
	if err != nil {
		return response.Error(err)
	}

	paginator.SetURL(c.OriginalURL())

	return response.View("user/index", helper.DataMap{
		"users":     users,
		"paginator": paginator,
	})
}

// Create render user create form
func (ctl UserController) Create(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	return response.View("user/create", helper.DataMap{
		"content_footer": true,
	})
}

// Store user store process
func (ctl UserController) Store(c *fiber.Ctx) error {
	var form form.User
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(true, 0, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Store(object.NewUserInput(
		form.Name,
		form.Email,
		form.Password,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Redirect(c, "system/user")
}

// Edit render user edit form
func (ctl UserController) Edit(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	user, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return response.ErrorWithCode(findErr, fiber.StatusNotFound)
	}

	return response.View("user/edit", helper.DataMap{
		"user":           user,
		"content_footer": true,
	})
}

// Update user update process
func (ctl UserController) Update(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	var form form.User

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(false, uint(id), ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Update(uint(id), object.NewUserInput(
		form.Name,
		form.Email,
		form.Password,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message(), Messages{})
	return response.Back(c)
}

// Delete user delete process
func (ctl UserController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message(), Messages{})
	return response.Redirect(c, "system/user")
}
