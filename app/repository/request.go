package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
	"gorm.io/gorm"
)

// RequestRepository database structure
type RequestRepository struct {
	db pkg.Database
}

// NewRequestRepository creates a new request repository
func NewRequestRepository(db pkg.Database) RequestRepository {
	return RequestRepository{
		db: db,
	}
}

// GetAll gets all requests
func (r RequestRepository) GetAll() (requests []model.Request, err error) {
	return requests, r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User").
		Find(&requests).
		Error
}

// Save request
func (r RequestRepository) Save(request model.Request) (model.Request, error) {
	return request, r.db.GormDB.Create(&request).Error
}

// SaveWithTagIDs request and request_tag
func (r RequestRepository) SaveWithTagIDs(request model.Request, tagIDs []uint) (model.Request, error) {
	if err := r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		if requestErr := tx.Create(&request).Error; requestErr != nil {
			return requestErr
		}

		var requestTags = make([]model.RequestTag, len(tagIDs))
		for index, id := range tagIDs {
			requestTags[index] = model.RequestTag{
				RequestID: request.ID,
				TagID:     id,
			}
		}
		return tx.Create(&requestTags).Error
	}); err != nil {
		return request, err
	}
	return request, nil
}

// Update updates request
func (r RequestRepository) Update(request model.Request) (model.Request, error) {
	return request, r.db.GormDB.Save(&request).Error
}

// UpdateWithTagIDs request and request_tag
func (r RequestRepository) UpdateWithTagIDs(request model.Request, tagIDs []uint) (model.Request, error) {
	if err := r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		if deleteErr := tx.Where("request_id = ?", request.ID).Delete(&model.RequestTag{}).Error; deleteErr != nil {
			return deleteErr
		}

		if requestErr := tx.Omit("Tags").Save(&request).Error; requestErr != nil {
			return requestErr
		}

		var requestTags = make([]model.RequestTag, len(tagIDs))
		for index, id := range tagIDs {
			requestTags[index] = model.RequestTag{
				RequestID: request.ID,
				TagID:     id,
			}
		}
		return tx.Create(&requestTags).Error
	}); err != nil {
		return request, err
	}
	return request, nil
}

// GetOne gets ont request
func (r RequestRepository) GetOne(id uint) (request model.Request, err error) {
	return request, r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User").
		Where("id = ?", id).
		First(&request).
		Error
}

// GetOneWithSuggests gets ont request
func (r RequestRepository) GetOneWithSuggests(id uint) (request model.Request, err error) {
	return request, r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User").
		Preload("Suggests.Suggester").
		Where("id = ?", id).
		First(&request).
		Error
}

// Delete deletes the row of data
func (r RequestRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Request{}).Error
}
