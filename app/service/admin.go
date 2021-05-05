package service

import (
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminService service logic
type AdminService struct {
	repository repository.AdminRepository
	logger     pkg.Logger
}

// NewAdminService new service
func NewAdminService(
	repository repository.AdminRepository,
	logger pkg.Logger,
) AdminService {
	return AdminService{
		repository: repository,
		logger:     logger,
	}
}

// Search search admins
func (s AdminService) Search() ([]model.Admin, error) {
	return s.repository.GetAll()
}

// Store create admin
func (s AdminService) Store(form form.Admin) (model.Admin, error) {
	admin := model.Admin{
		Name:  form.Name,
		Email: form.Email,
		Pass:  []byte(form.Password),
	}
	return s.repository.Save(admin)
}

// Update edit admin
func (s AdminService) Update(id uint, form form.Admin) (model.Admin, error) {
	admin, err := s.Find(id)
	if err != nil {
		return admin, err
	}

	admin.Name = form.Name
	admin.Email = form.Email
	admin.Pass = []byte(form.Password)

	return s.repository.Update(admin)
}

// Find get admin
func (s AdminService) Find(id uint) (admin model.Admin, err error) {
	return s.repository.GetOne(id)
}

// Delete remove admin
func (s AdminService) Delete(id uint) error {
	return s.repository.Delete(id)
}
