package service

import (
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// TodoService service logic
type TodoService struct {
	Repository repository.TodoRepository
	Query      query.TodoQuery
	logger     pkg.Logger
}

// NewTodoService new service
func NewTodoService(
	repository repository.TodoRepository,
	query query.TodoQuery,
	logger pkg.Logger,
) TodoService {
	return TodoService{
		Repository: repository,
		Query:      query,
		logger:     logger,
	}
}

// Search search todos
func (s TodoService) Search(form form.TodoSearch, limit int) ([]model.Todo, error) {
	return s.Query.Search(form, limit)
}

// Store create todo
func (s TodoService) Store(form form.Todo, adminID uint) (model.Todo, error) {

	todo := model.Todo{
		Text:    form.Text,
		Status:  model.TodoStatusFromString(form.Status),
		AdminID: adminID,
	}
	return s.Repository.Save(todo)
}

// Update edit todo
func (s TodoService) Update(id uint, form form.Todo) (model.Todo, error) {
	todo, err := s.Find(id)
	if err != nil {
		return todo, err
	}

	todo.Text = form.Text
	todo.Status = model.TodoStatusFromString(form.Status)

	return s.Repository.Update(todo)
}

// ChangeStatus todo status
func (s TodoService) ChangeStatus(id uint, status string) (model.Todo, error) {
	todo, err := s.Find(id)
	if err != nil {
		return todo, err
	}

	todo.Status = model.TodoStatusFromString(status)

	return s.Repository.Update(todo)
}

// Find get todo
func (s TodoService) Find(id uint) (todo model.Todo, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove todo
func (s TodoService) Delete(id uint) error {
	return s.Repository.Delete(id)
}
