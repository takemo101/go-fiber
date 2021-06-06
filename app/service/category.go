package service

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// CategoryService service logic
type CategoryService struct {
	Repository repository.CategoryRepository
	logger     pkg.Logger
}

// NewCategoryService new service
func NewCategoryService(
	repository repository.CategoryRepository,
	logger pkg.Logger,
) CategoryService {
	return CategoryService{
		Repository: repository,
		logger:     logger,
	}
}

// GetAll categorys
func (s CategoryService) FindAll() ([]model.Category, error) {
	return s.Repository.GetAll()
}

// Store create category
func (s CategoryService) Store(object object.CategoryInput) (model.Category, error) {

	max := s.Repository.MaxSort()
	category := model.Category{
		Name:     object.GetName(),
		IsActive: object.GetIsActive(),
		Sort:     max + 1,
	}
	return s.Repository.Save(category)
}

// Update edit category
func (s CategoryService) Update(id uint, object object.CategoryInput) (model.Category, error) {
	category, err := s.Find(id)
	if err != nil {
		return category, err
	}

	category.Name = object.GetName()
	category.IsActive = object.GetIsActive()

	return s.Repository.Update(category)
}

// Find get category
func (s CategoryService) Find(id uint) (category model.Category, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove category
func (s CategoryService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// Sort category sort by ids
func (s CategoryService) Sort(object object.CategorySortInput) error {
	return s.Repository.SortUpdate(object.GetIDs())
}
