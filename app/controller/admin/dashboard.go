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
	logger          pkg.Logger
	config          pkg.Config
	value           support.RequestValue
	todoQuery       query.TodoQuery
	requestQuery    query.RequestQuery
	discussionQuery query.DiscussionQuery
}

// NewDashboardController is create dashboard
func NewDashboardController(
	logger pkg.Logger,
	config pkg.Config,
	value support.RequestValue,
	todoQuery query.TodoQuery,
	requestQuery query.RequestQuery,
	discussionQuery query.DiscussionQuery,
) DashboardController {
	return DashboardController{
		logger:          logger,
		config:          config,
		value:           value,
		todoQuery:       todoQuery,
		requestQuery:    requestQuery,
		discussionQuery: discussionQuery,
	}
}

// Dashboard render home
func (ctl DashboardController) Dashboard(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	todos, todoErr := ctl.todoQuery.GetUpdateTodos(10)
	if todoErr != nil {
		return response.Error(todoErr)
	}

	requests, requestErr := ctl.requestQuery.GetUpdateRequests(10)
	if requestErr != nil {
		return response.Error(requestErr)
	}

	discussions, discussionErr := ctl.discussionQuery.GetUpdateDiscussions(10)
	if discussionErr != nil {
		return response.Error(discussionErr)
	}

	return response.View("home", helper.DataMap{
		"todos":       todos,
		"requests":    requests,
		"discussions": discussions,
		"config":      ctl.config,
	})
}
