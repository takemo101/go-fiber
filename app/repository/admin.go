package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminRepository database structure
type AdminRepository struct {
	db     pkg.Database
	logger pkg.Logger
}

// NewAdminRepository creates a new admin repository
func NewAdminRepository(db pkg.Database, logger pkg.Logger) AdminRepository {
	return AdminRepository{
		db:     db,
		logger: logger,
	}
}

// GetAll gets all admins
func (r AdminRepository) GetAll() (admins []model.Admin, err error) {
	return admins, r.db.GormDB.Find(&admins).Error
}

// Save admin
func (r AdminRepository) Save(admin model.Admin) (model.Admin, error) {
	return admin, r.db.GormDB.Create(&admin).Error
}

// Update updates admin
func (r AdminRepository) Update(admin model.Admin) (model.Admin, error) {
	return admin, r.db.GormDB.Save(&admin).Error
}

// GetOne gets ont admin
func (r AdminRepository) GetOne(id uint) (admin model.Admin, err error) {
	return admin, r.db.GormDB.Where("id = ?", id).First(&admin).Error
}

// Delete deletes the row of data
func (r AdminRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Admin{}).Error
}

// FindByName is find by name
func (r AdminRepository) FindByName(name string) (admin model.Admin, err error) {
	return admin, r.db.GormDB.Where("name = ?", name).First(&admin).Error
}
