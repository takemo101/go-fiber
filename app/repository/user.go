package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// UserRepository database structure
type UserRepository struct {
	db     pkg.Database
	logger pkg.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db pkg.Database, logger pkg.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// GetAll gets all users
func (r UserRepository) GetAll() (users []model.User, err error) {
	return users, r.db.GormDB.Find(&users).Error
}

// Save user
func (r UserRepository) Save(user model.User) (model.User, error) {
	return user, r.db.GormDB.Create(&user).Error
}

// Update updates user
func (r UserRepository) Update(user model.User) (model.User, error) {
	return user, r.db.GormDB.Save(&user).Error
}

// GetOne gets ont user
func (r UserRepository) GetOne(id uint) (user model.User, err error) {
	return user, r.db.GormDB.Where("id = ?", id).First(&user).Error
}

// Delete deletes the row of data
func (r UserRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.User{}).Error
}

// FindByEmail is find by email
func (r UserRepository) FindByEmail(email string) (user model.User, err error) {
	return user, r.db.GormDB.Where("email = ?", email).First(&user).Error
}
