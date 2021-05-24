package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/pkg"
)

// DashboardController is home dashboard
type DashboardController struct {
	logger    pkg.Logger
	config    pkg.Config
	response  *helper.ResponseHelper
	todoQuery query.TodoQuery
}

// NewDashboardController is create dashboard
func NewDashboardController(
	logger pkg.Logger,
	config pkg.Config,
	response *helper.ResponseHelper,
	todoQuery query.TodoQuery,
) DashboardController {
	return DashboardController{
		logger:    logger,
		config:    config,
		response:  response,
		todoQuery: todoQuery,
	}
}

// Dashboard render home
func (ctl DashboardController) Dashboard(c *fiber.Ctx) error {

	todos, todoErr := ctl.todoQuery.GetUpdateTodos(10)
	if todoErr != nil {
		return ctl.response.Error(todoErr)
	}

	return ctl.response.View("home", helper.DataMap{
		"todos":  todos,
		"config": ctl.config,
	})
}
