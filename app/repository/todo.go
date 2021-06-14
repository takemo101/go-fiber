package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// TodoRepository database structure
type TodoRepository struct {
	db pkg.Database
}

// NewTodoRepository creates a new todo repository
func NewTodoRepository(db pkg.Database) TodoRepository {
	return TodoRepository{
		db: db,
	}
}

// GetAll get all todos
func (r TodoRepository) GetAll() (todos []model.Todo, err error) {
	return todos, r.db.GormDB.Preload("Admin").Find(&todos).Error
}

// Save todo
func (r TodoRepository) Save(todo model.Todo) (model.Todo, error) {
	return todo, r.db.GormDB.Create(&todo).Error
}

// Update updates todo
func (r TodoRepository) Update(todo model.Todo) (model.Todo, error) {
	return todo, r.db.GormDB.Save(&todo).Error
}

// GetOne get one todo
func (r TodoRepository) GetOne(id uint) (todo model.Todo, err error) {
	return todo, r.db.GormDB.Where("id = ?", id).First(&todo).Error
}

// Delete deletes the row of data
func (r TodoRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Todo{}).Error
}
