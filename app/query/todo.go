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
func (r TodoQuery) Search(object object.TodoSearchInput, limit int) (todos []model.Todo, paginator Paginator, err error) {
	db := r.db.GormDB.Preload("Admin")

	err = Paging(&PagingParameter{
		DB:      db,
		Page:    object.GetPage(),
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &todos, &paginator)

	return todos, paginator, err
}

// SearchYour get todos
func (r TodoQuery) SearchYour(object object.TodoSearchInput, adminID uint, limit int) (todos []model.Todo, paginator Paginator, err error) {
	db := r.db.GormDB.
		Preload("Admin").
		Where("admin_id = ?", adminID)

	err = Paging(&PagingParameter{
		DB:      db,
		Page:    object.GetPage(),
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &todos, &paginator)

	return todos, paginator, err
}

// GetUpdateTodos get todos order by update_at
func (r TodoQuery) GetUpdateTodos(limit int) (todos []model.Todo, err error) {
	return todos, r.db.GormDB.Preload("Admin").Order("updated_at desc").Limit(limit).Find(&todos).Error
}
