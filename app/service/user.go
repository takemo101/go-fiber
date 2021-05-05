package service

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// UserService service logic
type UserService struct {
	repository repository.UserRepository
	logger     pkg.Logger
}

// NewUserService new service
func NewUserService(
	repository repository.UserRepository,
	logger pkg.Logger,
) UserService {
	return UserService{
		repository: repository,
		logger:     logger,
	}
}

// Search search users
func (s UserService) Search() (users []model.User, err error) {
	return s.repository.GetAll()
}

// Store create user
func (s UserService) Store(user model.User) (model.User, error) {
	return s.repository.Save(user)
}

// Update edit user
func (s UserService) Update(user model.User) (model.User, error) {
	return s.repository.Update(user)
}

// Find get user
func (s UserService) Find(id uint) (user model.User, err error) {
	return s.repository.GetOne(id)
}

// Delete remove user
func (s UserService) Delete(id uint) error {
	return s.repository.Delete(id)
}
