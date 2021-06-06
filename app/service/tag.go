package service

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// TagService service logic
type TagService struct {
	Repository repository.TagRepository
	logger     pkg.Logger
}

// NewTagService new service
func NewTagService(
	repository repository.TagRepository,
	logger pkg.Logger,
) TagService {
	return TagService{
		Repository: repository,
		logger:     logger,
	}
}

// GetAll tags
func (s TagService) FindAll() ([]model.Tag, error) {
	return s.Repository.GetAll()
}

// Store create tag
func (s TagService) Store(object object.TagInput) (model.Tag, error) {

	max := s.Repository.MaxSort()
	tag := model.Tag{
		Name: object.GetName(),
		Sort: max + 1,
	}
	return s.Repository.Save(tag)
}

// Update edit tag
func (s TagService) Update(id uint, object object.TagInput) (model.Tag, error) {
	tag, err := s.Find(id)
	if err != nil {
		return tag, err
	}

	tag.Name = object.GetName()

	return s.Repository.Update(tag)
}

// Find get tag
func (s TagService) Find(id uint) (tag model.Tag, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove tag
func (s TagService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// Sort tag sort by ids
func (s TagService) Sort(object object.TagSortInput) error {
	return s.Repository.SortUpdate(object.GetIDs())
}
