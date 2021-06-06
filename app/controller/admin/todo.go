package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// TodoController is todo
type TodoController struct {
	logger  pkg.Logger
	service service.TodoService
	value   support.RequestValue
}

// NewTodoController is create todo controller
func NewTodoController(
	logger pkg.Logger,
	service service.TodoService,
	value support.RequestValue,
) TodoController {
	return TodoController{
		logger:  logger,
		service: service,
		value:   value,
	}
}

// Index render todo list
func (ctl TodoController) Index(c *fiber.Ctx) error {
	var form form.TodoSearch
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	todos, err := ctl.service.Search(object.NewTodoSearchInput(
		form.Keyword,
		form.Page,
	), 20)
	if err != nil {
		return response.Error(err)
	}

	response.JS(helper.DataMap{
		"statuses": model.ToTodoStatusArray(),
	})
	return response.View("todo/index", helper.DataMap{
		"todos":    todos,
		"statuses": model.ToTodoStatusArray(),
	})
}

// Your render todo list
func (ctl TodoController) Your(c *fiber.Ctx) error {
	var form form.TodoSearch
	auth := ctl.value.GetSessionAdminAuth(c)
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	todos, err := ctl.service.SearchYour(object.NewTodoSearchInput(
		form.Keyword,
		form.Page,
	), auth.ID(), 20)
	if err != nil {
		return response.Error(err)
	}

	response.JS(helper.DataMap{
		"statuses": model.ToTodoStatusArray(),
	})
	return response.View("todo/your", helper.DataMap{
		"todos":    todos,
		"statuses": model.ToTodoStatusArray(),
	})
}

// Store todo store process
func (ctl TodoController) Store(c *fiber.Ctx) error {
	var form form.Todo
	auth := ctl.value.GetSessionAdminAuth(c)
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Store(object.NewTodoInput(
		form.Text,
		form.Status,
	), auth.ID()); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Back(c)
}

// ChangeStatus todo update process
func (ctl TodoController) ChangeStatus(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.JsonError(c, convErr)
	}

	uID := uint(id)

	// check todo owner
	if err := ctl.checkTodoOwner(c, uID); err != nil {
		return response.JsonError(c, err)
	}

	var form form.Todo

	if err := c.BodyParser(&form); err != nil {
		return response.JsonError(c, err)
	}

	if _, err := ctl.service.ChangeStatus(uID, form.Status); err != nil {
		return response.JsonError(c, err)
	}

	return response.JsonSuccess(c, "change stauts successfully")
}

// Delete todo delete process
func (ctl TodoController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	uID := uint(id)

	// check todo owner
	if err := ctl.checkTodoOwner(c, uID); err != nil {
		return response.Error(err)
	}

	if err := ctl.service.Delete(uID); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message(), Messages{})
	return response.Back(c)
}

// checkTodoOwner todo admin owner check
func (ctl TodoController) checkTodoOwner(c *fiber.Ctx, id uint) error {
	// find todo
	todo, findErr := ctl.service.Find(id)
	if findErr != nil {
		return findErr
	}

	auth := ctl.value.GetSessionAdminAuth(c)

	// todo owner check
	if !ctl.service.CheckOwner(todo, *auth.Admin()) {
		return fiber.ErrUnauthorized
	}
	return nil
}
