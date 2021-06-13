package service

import (
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// UserService service logic
type UserService struct {
	Repository repository.UserRepository
	Query      query.UserQuery
	logger     pkg.Logger
}

// NewUserService new service
func NewUserService(
	repository repository.UserRepository,
	query query.UserQuery,
	logger pkg.Logger,
) UserService {
	return UserService{
		Repository: repository,
		Query:      query,
		logger:     logger,
	}
}

// Search search users
func (s UserService) Search(object object.UserSearchInput, limit int) ([]model.User, query.Paginator, error) {
	return s.Query.Search(object, limit)
}

// Store create user
func (s UserService) Store(object object.UserInput) (model.User, error) {
	pass, passErr := s.GeneratePass(object.GetPass())
	if passErr != nil {
		return model.User{}, passErr
	}

	user := model.User{
		Name:  object.GetName(),
		Email: object.GetEmail(),
		Pass:  pass,
	}
	return s.Repository.Save(user)
}

// StoreByModel create user by model
func (s UserService) StoreByModel(user model.User) (model.User, error) {
	pass, passErr := s.GeneratePass(user.Pass)
	if passErr != nil {
		return model.User{}, passErr
	}

	user.Pass = pass
	return s.Repository.Save(user)
}

// Update edit user
func (s UserService) Update(id uint, object object.UserInput) (model.User, error) {
	user, err := s.Find(id)
	if err != nil {
		return user, err
	}

	user.Name = object.GetName()
	user.Email = object.GetEmail()
	if object.HasPass() {
		pass, passErr := s.GeneratePass(object.GetPass())
		if passErr != nil {
			return model.User{}, passErr
		}
		user.Pass = pass
	}
	return s.Repository.Update(user)
}

// Find get user
func (s UserService) Find(id uint) (user model.User, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove user
func (s UserService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// GeneratePass generate hash password
func (s UserService) GeneratePass(pass []byte) ([]byte, error) {
	hash, err := helper.CreatePass(pass)
	if err != nil {
		return nil, err
	}
	return []byte(hash), nil
}
