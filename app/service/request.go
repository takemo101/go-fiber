package service

import (
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// RequestService service logic
type RequestService struct {
	Repository repository.RequestRepository
	Query      query.RequestQuery
	file       helper.FileHelper
	logger     pkg.Logger
}

// NewRequestService new service
func NewRequestService(
	repository repository.RequestRepository,
	query query.RequestQuery,
	file helper.FileHelper,
	logger pkg.Logger,
) RequestService {
	return RequestService{
		Repository: repository,
		Query:      query,
		file:       file,
		logger:     logger,
	}
}

// Search search requests
func (s RequestService) Search(object object.RequestSearchInput, limit int) ([]model.Request, error) {
	return s.Query.Search(object, limit)
}

// Store create request
func (s RequestService) Store(object object.RequestInput, userID uint) (model.Request, error) {

	request := model.Request{
		Title:      object.GetTitle(),
		Content:    object.GetContent(),
		Thumbnail:  object.GetThumbnail(),
		Status:     object.GetStatus(),
		CategoryID: object.GetCategoryID(),
		UserID:     userID,
	}
	return s.Repository.SaveWithTagIDs(request, object.GetTagIDs())
}

// Update edit request
func (s RequestService) Update(id uint, object object.RequestInput) (model.Request, error) {
	request, err := s.Find(id)
	if err != nil {
		return request, err
	}

	request.Title = object.GetTitle()
	request.Content = object.GetContent()
	request.Status = object.GetStatus()
	request.CategoryID = object.GetCategoryID()
	if object.HasThumbnail() {
		s.file.RemovePublic(request.Thumbnail)
		request.Thumbnail = object.GetThumbnail()
	}

	return s.Repository.UpdateWithTagIDs(request, object.GetTagIDs())
}

// ChangeStatus request status
func (s RequestService) ChangeStatus(id uint, status string) (model.Request, error) {
	request, err := s.Find(id)
	if err != nil {
		return request, err
	}

	request.Status = model.RequestStatus(status)

	return s.Repository.Update(request)
}

// Find get request
func (s RequestService) Find(id uint) (request model.Request, err error) {
	return s.Repository.GetOne(id)
}

// Find get request with suggests
func (s RequestService) FindWithSuggests(id uint) (request model.Request, err error) {
	return s.Repository.GetOneWithSuggests(id)
}

// Delete remove request
func (s RequestService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// CheckOwner check request owner
func (s RequestService) CheckOwner(request model.Request, user model.User) bool {
	return request.UserID == user.ID
}
