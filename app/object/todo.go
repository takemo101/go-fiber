package object

import (
	"strconv"
	"strings"

	"github.com/takemo101/go-fiber/app/model"
)

// TodoInput input form to service
type TodoInput struct {
	text   string
	status string
}

func NewTodoInput(
	text string,
	status string,
) TodoInput {
	return TodoInput{
		text:   text,
		status: status,
	}
}

func (o TodoInput) GetText() string {
	return o.text
}

func (o TodoInput) GetStatus() model.TodoStatus {
	return model.TodoStatus(o.status)
}

// TodoSearchInput search form to service
type TodoSearchInput struct {
	keyword string
	page    string
}

func NewTodoSearchInput(
	keyword string,
	page string,
) TodoSearchInput {
	return TodoSearchInput{
		keyword: keyword,
		page:    page,
	}
}

func (o TodoSearchInput) GetKeyword() string {
	return strings.TrimSpace(o.keyword)
}

func (o TodoSearchInput) GetPage() int {
	if page, err := strconv.Atoi(o.page); err == nil {
		return page
	}
	return 0
}
