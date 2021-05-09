package service

import (
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/query"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminService service logic
type AdminService struct {
	Repository repository.AdminRepository
	Query      query.AdminQuery
	logger     pkg.Logger
}

// NewAdminService new service
func NewAdminService(
	repository repository.AdminRepository,
	query query.AdminQuery,
	logger pkg.Logger,
) AdminService {
	return AdminService{
		Repository: repository,
		Query:      query,
		logger:     logger,
	}
}

// Search search admins
func (s AdminService) Search(form form.AdminSearch, limit int) ([]model.Admin, error) {
	return s.Query.Search(form, limit)
}

// Store create admin
func (s AdminService) Store(form form.Admin) (model.Admin, error) {
	pass, passErr := s.GeneratePass(form.Password)
	if passErr != nil {
		return model.Admin{}, passErr
	}

	admin := model.Admin{
		Name:  form.Name,
		Email: form.Email,
		Pass:  pass,
		Role:  model.RoleFromString(form.Role),
	}
	return s.Repository.Save(admin)
}

// Update edit admin
func (s AdminService) Update(id uint, form form.Admin) (model.Admin, error) {
	admin, err := s.Find(id)
	if err != nil {
		return admin, err
	}

	admin.Name = form.Name
	admin.Email = form.Email

	if form.Role != "" {
		admin.Role = model.RoleFromString(form.Role)
	}

	if form.Password != "" {
		pass, passErr := s.GeneratePass(form.Password)
		if passErr != nil {
			return model.Admin{}, passErr
		}
		admin.Pass = pass
	}

	return s.Repository.Update(admin)
}

// Find get admin
func (s AdminService) Find(id uint) (admin model.Admin, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove admin
func (s AdminService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// GeneratePass generate hash password
func (s AdminService) GeneratePass(pass string) ([]byte, error) {
	hash, err := helper.CreatePass(pass)
	if err != nil {
		return nil, err
	}
	return []byte(hash), nil
}
