package service

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
)

// DiscussionService service logic
type DiscussionService struct {
	Repository        repository.DiscussionRepository
	SuggestRepository repository.SuggestRepository
	Query             query.DiscussionQuery
}

// NewDiscussionService new service
func NewDiscussionService(
	repository repository.DiscussionRepository,
	suggestRepository repository.SuggestRepository,
	query query.DiscussionQuery,
) DiscussionService {
	return DiscussionService{
		Repository:        repository,
		SuggestRepository: suggestRepository,
		Query:             query,
	}
}

// Search search requests
func (s DiscussionService) Search(object object.DiscussionSearchInput, limit int) ([]model.Discussion, query.Paginator, error) {
	return s.Query.Search(object, limit)
}

// GetAll discussions
func (s DiscussionService) FindAll() ([]model.Discussion, error) {
	return s.Repository.GetAll()
}

// Find get discussion
func (s DiscussionService) Find(id uint) (discussion model.Discussion, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove discussion
func (s DiscussionService) Delete(id uint) error {
	return s.Repository.Delete(id)
}
