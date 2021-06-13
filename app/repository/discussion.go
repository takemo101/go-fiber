package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
	"gorm.io/gorm"
)

// DiscussionRepository database structure
type DiscussionRepository struct {
	db pkg.Database
}

// NewDiscussionRepository creates a new discussion repository
func NewDiscussionRepository(db pkg.Database) DiscussionRepository {
	return DiscussionRepository{
		db: db,
	}
}

// GetAll gets all discussions
func (r DiscussionRepository) GetAll() (discussions []model.Discussion, err error) {
	return discussions, r.db.GormDB.Order("sort asc").Find(&discussions).Error
}

// Save discussion
func (r DiscussionRepository) Save(discussion model.Discussion) (model.Discussion, error) {
	return discussion, r.db.GormDB.Create(&discussion).Error
}

// Save discussion
func (r DiscussionRepository) SaveWithUpdateSuggest(discussion model.Discussion, suggest model.Suggest) (model.Discussion, error) {
	if err := r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		if createErr := tx.Create(&discussion).Error; createErr != nil {
			return createErr
		}
		return tx.Omit("Request").Omit("Suggester").Save(&suggest).Error
	}); err != nil {
		return discussion, err
	}
	return discussion, nil
}

// Update updates discussion
func (r DiscussionRepository) Update(discussion model.Discussion) (model.Discussion, error) {
	return discussion, r.db.GormDB.Save(&discussion).Error
}

// GetOne gets ont discussion
func (r DiscussionRepository) GetOne(id uint) (discussion model.Discussion, err error) {
	return discussion, r.db.GormDB.Where("id = ?", id).First(&discussion).Error
}

// Delete deletes the row of data
func (r DiscussionRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Discussion{}).Error
}
