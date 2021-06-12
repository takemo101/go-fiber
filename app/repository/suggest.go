package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
	"gorm.io/gorm"
)

// SuggestRepository database structure
type SuggestRepository struct {
	db     pkg.Database
	logger pkg.Logger
}

// NewSuggestRepository creates a new suggest repository
func NewSuggestRepository(db pkg.Database, logger pkg.Logger) SuggestRepository {
	return SuggestRepository{
		db:     db,
		logger: logger,
	}
}

// GetAll gets all suggests
func (r SuggestRepository) GetAll() (suggests []model.Suggest, err error) {
	return suggests, r.db.GormDB.Order("sort asc").Find(&suggests).Error
}

// Save suggest
func (r SuggestRepository) Save(suggest model.Suggest) (model.Suggest, error) {
	return suggest, r.db.GormDB.Create(&suggest).Error
}

// Save suggest
func (r SuggestRepository) SaveWithDiscussion(suggest model.Suggest, discussion model.Discussion) (model.Suggest, error) {
	if err := r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		if createErr := tx.Create(&suggest).Error; createErr != nil {
			return createErr
		}
		discussion.SuggestID = suggest.ID
		return tx.Omit("Suggest").Create(&discussion).Error
	}); err != nil {
		return suggest, err
	}
	return suggest, nil
}

// Update updates suggest
func (r SuggestRepository) Update(suggest model.Suggest) (model.Suggest, error) {
	return suggest, r.db.GormDB.Save(&suggest).Error
}

// GetOne gets ont suggest
func (r SuggestRepository) GetOne(id uint) (suggest model.Suggest, err error) {
	return suggest, r.db.GormDB.Where("id = ?", id).Preload("Suggester").Preload("Request.User").First(&suggest).Error
}

// GetOneWithDiscussions gets ont suggest
func (r SuggestRepository) GetOneWithDiscussions(id uint) (suggest model.Suggest, err error) {
	return suggest, r.db.GormDB.Where("id = ?", id).Preload("Suggester").Preload("Request.User").Preload("Discussions").First(&suggest).Error
}

// Delete deletes the row of data
func (r SuggestRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Suggest{}).Error
}

// ExistsByRequestIDAndSuggesterID is exists by requestID and UserID
func (r SuggestRepository) ExistsByRequestIDAndSuggesterID(requestID uint, suggesterID uint) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Suggest{}).
		Where("request_id = ?", requestID).
		Where("suggester_id = ?", suggesterID).
		Count(&count).
		Error

	return (count > 0), err
}
