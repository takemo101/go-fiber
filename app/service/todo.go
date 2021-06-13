package service

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
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
func (s TodoService) Search(object object.TodoSearchInput, limit int) ([]model.Todo, query.Paginator, error) {
	return s.Query.Search(object, limit)
}

// SearchYour search todos
func (s TodoService) SearchYour(object object.TodoSearchInput, adminID uint, limit int) ([]model.Todo, query.Paginator, error) {
	return s.Query.SearchYour(object, adminID, limit)
}

// Store create todo
func (s TodoService) Store(object object.TodoInput, adminID uint) (model.Todo, error) {

	todo := model.Todo{
		Text:    object.GetText(),
		Status:  object.GetStatus(),
		AdminID: adminID,
	}
	return s.Repository.Save(todo)
}

// Update edit todo
func (s TodoService) Update(id uint, object object.TodoInput) (model.Todo, error) {
	todo, err := s.Find(id)
	if err != nil {
		return todo, err
	}

	todo.Text = object.GetText()
	todo.Status = object.GetStatus()

	return s.Repository.Update(todo)
}

// ChangeStatus todo status
func (s TodoService) ChangeStatus(id uint, status string) (model.Todo, error) {
	todo, err := s.Find(id)
	if err != nil {
		return todo, err
	}

	todo.Status = model.TodoStatus(status)

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

// CheckOwner check todo owner
func (s TodoService) CheckOwner(todo model.Todo, admin model.Admin) bool {
	return todo.AdminID == admin.ID
}
