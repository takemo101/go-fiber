package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// TodoController is todo
type TodoController struct {
	logger   pkg.Logger
	response *helper.ResponseHelper
	service  service.TodoService
	auth     middleware.SessionAdminAuth
}

// NewTodoController is create todo controller
func NewTodoController(
	logger pkg.Logger,
	response *helper.ResponseHelper,
	service service.TodoService,
	auth middleware.SessionAdminAuth,
) TodoController {
	return TodoController{
		logger:   logger,
		response: response,
		service:  service,
		auth:     auth,
	}
}

// Index render todo list
func (ctl TodoController) Index(c *fiber.Ctx) error {
	var form form.TodoSearch

	if err := c.QueryParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	todos, err := ctl.service.Search(form, 20)
	if err != nil {
		return ctl.response.Error(err)
	}

	ctl.response.JS(helper.DataMap{
		"statuses": model.ToTodoStatusArray(),
	})
	return ctl.response.View("todo/index", helper.DataMap{
		"todos":    todos,
		"statuses": model.ToTodoStatusArray(),
	})
}

// Your render todo list
func (ctl TodoController) Your(c *fiber.Ctx) error {
	var form form.TodoSearch

	if err := c.QueryParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	todos, err := ctl.service.SearchYour(form, ctl.auth.Auth.ID(), 20)
	if err != nil {
		return ctl.response.Error(err)
	}

	ctl.response.JS(helper.DataMap{
		"statuses": model.ToTodoStatusArray(),
	})
	return ctl.response.View("todo/your", helper.DataMap{
		"todos":    todos,
		"statuses": model.ToTodoStatusArray(),
	})
}

// Store todo store process
func (ctl TodoController) Store(c *fiber.Ctx) error {
	var form form.Todo

	if err := c.BodyParser(&form); err != nil {
		return ctl.response.Error(err)
	}

	if err := form.Validate(); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return ctl.response.Back(c)
	}

	if _, err := ctl.service.Store(form, ctl.auth.Auth.ID()); err != nil {
		return ctl.response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message())
	return ctl.response.Back(c)
}

// ChangeStatus todo update process
func (ctl TodoController) ChangeStatus(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.response.JsonError(c, convErr)
	}

	uID := uint(id)

	// check todo owner
	if err := ctl.checkTodoOwner(uID); err != nil {
		return ctl.response.JsonError(c, err)
	}

	var form form.Todo

	if err := c.BodyParser(&form); err != nil {
		return ctl.response.JsonError(c, err)
	}

	if _, err := ctl.service.ChangeStatus(uID, form.Status); err != nil {
		return ctl.response.JsonError(c, err)
	}

	return ctl.response.JsonSuccess(c, "change stauts successfully")
}

// Delete todo delete process
func (ctl TodoController) Delete(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.response.Error(convErr)
	}

	uID := uint(id)

	// check todo owner
	if err := ctl.checkTodoOwner(uID); err != nil {
		return ctl.response.Error(err)
	}

	if err := ctl.service.Delete(uID); err != nil {
		return ctl.response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message())
	return ctl.response.Back(c)
}

// checkTodoOwner todo admin owner check
func (ctl TodoController) checkTodoOwner(id uint) error {
	// find todo
	todo, findErr := ctl.service.Find(id)
	if findErr != nil {
		return findErr
	}

	// todo owner check
	if !ctl.service.CheckOwner(todo, *ctl.auth.Auth.Admin()) {
		return fiber.ErrUnauthorized
	}
	return nil
}
