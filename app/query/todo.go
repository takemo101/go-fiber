package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
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

// Search gettodos todos
func (r TodoQuery) Search(object object.TodoSearchInput, limit int) (todos []model.Todo, err error) {
	return todos, r.db.GormDB.Preload("Admin").
		Order("id desc").
		Limit(limit).
		Find(&todos).Error
}

// SearchYour get todos
func (r TodoQuery) SearchYour(object object.TodoSearchInput, adminID uint, limit int) (todos []model.Todo, err error) {
	return todos, r.db.GormDB.Preload("Admin").
		Where("admin_id = ?", adminID).
		Order("id desc").
		Limit(limit).
		Find(&todos).Error
}

// GetUpdateMenus get todos order by update_at
func (r TodoQuery) GetUpdateTodos(limit int) (todos []model.Todo, err error) {
	return todos, r.db.GormDB.Preload("Admin").Order("updated_at desc").Limit(limit).Find(&todos).Error
}
