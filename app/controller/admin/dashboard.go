package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// DashboardController is home dashboard
type DashboardController struct {
	logger    pkg.Logger
	config    pkg.Config
	value     support.RequestValue
	todoQuery query.TodoQuery
	menuQuery query.MenuQuery
}

// NewDashboardController is create dashboard
func NewDashboardController(
	logger pkg.Logger,
	config pkg.Config,
	value support.RequestValue,
	todoQuery query.TodoQuery,
	menuQuery query.MenuQuery,
) DashboardController {
	return DashboardController{
		logger:    logger,
		config:    config,
		value:     value,
		todoQuery: todoQuery,
		menuQuery: menuQuery,
	}
}

// Dashboard render home
func (ctl DashboardController) Dashboard(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	todos, todoErr := ctl.todoQuery.GetUpdateTodos(10)
	if todoErr != nil {
		return response.Error(todoErr)
	}

	menus, menuErr := ctl.menuQuery.GetUpdateMenus(10)
	if menuErr != nil {
		return response.Error(menuErr)
	}

	return response.View("home", helper.DataMap{
		"todos":  todos,
		"menus":  menus,
		"config": ctl.config,
	})
}
