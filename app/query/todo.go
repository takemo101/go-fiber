package query

import (
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// TodoQuery database structure
type TodoQuery struct {
	db pkg.Database
}

// NewTodoQuery creates a new todo query
func NewTodoQuery(db pkg.Database) TodoQuery {
	return TodoQuery{
		db: db,
	}
}

// Search gets limit todos
func (r TodoQuery) Search(form form.TodoSearch, limit int) (todos []model.Todo, err error) {
	return todos, r.db.GormDB.Preload("Admin").Order("id desc").Limit(limit).Find(&todos).Error
}
