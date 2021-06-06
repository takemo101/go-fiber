package service

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// MenuService service logic
type MenuService struct {
	Repository repository.MenuRepository
	Query      query.MenuQuery
	logger     pkg.Logger
}

// NewMenuService new service
func NewMenuService(
	repository repository.MenuRepository,
	query query.MenuQuery,
	logger pkg.Logger,
) MenuService {
	return MenuService{
		Repository: repository,
		Query:      query,
		logger:     logger,
	}
}

// Search search menus
func (s MenuService) Search(object object.MenuSearchInput, limit int) ([]model.Menu, error) {
	return s.Query.Search(object, limit)
}

// Store create menu
func (s MenuService) Store(object object.MenuInput, userID uint) (model.Menu, error) {

	menu := model.Menu{
		Title:      object.GetTitle(),
		Content:    object.GetContent(),
		Process:    object.GetProcess(),
		Status:     object.GetStatus(),
		CategoryID: object.GetCategoryID(),
		UserID:     userID,
	}
	return s.Repository.SaveWithTagIDs(menu, object.GetTagIDs())
}

// Update edit menu
func (s MenuService) Update(id uint, object object.MenuInput) (model.Menu, error) {
	menu, err := s.Find(id)
	if err != nil {
		return menu, err
	}

	menu.Title = object.GetTitle()
	menu.Content = object.GetContent()
	menu.Process = object.GetProcess()
	menu.Status = object.GetStatus()
	menu.CategoryID = object.GetCategoryID()

	return s.Repository.UpdateWithTagIDs(menu, object.GetTagIDs())
}

// ChangeStatus menu status
func (s MenuService) ChangeStatus(id uint, status string) (model.Menu, error) {
	menu, err := s.Find(id)
	if err != nil {
		return menu, err
	}

	menu.Status = model.MenuStatus(status)

	return s.Repository.Update(menu)
}

// Find get menu
func (s MenuService) Find(id uint) (menu model.Menu, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove menu
func (s MenuService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// CheckOwner check menu owner
func (s MenuService) CheckOwner(menu model.Menu, user model.User) bool {
	return menu.UserID == user.ID
}
