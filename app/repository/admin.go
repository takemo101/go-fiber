package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminRepository database structure
type AdminRepository struct {
	db pkg.Database
}

// NewAdminRepository creates a new admin repository
func NewAdminRepository(db pkg.Database) AdminRepository {
	return AdminRepository{
		db: db,
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

// GetOneByEmail is find by email
func (r AdminRepository) GetOneByEmail(email string) (admin model.Admin, err error) {
	return admin, r.db.GormDB.Where("email = ?", email).First(&admin).Error
}

// ExistsByEmail is exists by email
func (r AdminRepository) ExistsByEmail(email string) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Admin{}).
		Where("email = ?", email).
		Count(&count).
		Error

	return (count > 0), err
}

// ExistsByIDEmail is exists by id and email
func (r AdminRepository) ExistsByIDEmail(id uint, email string) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Admin{}).
		Where("id <> ? AND email = ?", id, email).
		Count(&count).
		Error

	return (count > 0), err
}
