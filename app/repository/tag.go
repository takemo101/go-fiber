package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
	"gorm.io/gorm"
)

// TagRepository database structure
type TagRepository struct {
	db pkg.Database
}

// NewTagRepository creates a new tag repository
func NewTagRepository(db pkg.Database) TagRepository {
	return TagRepository{
		db: db,
	}
}

// GetAll gets all tags
func (r TagRepository) GetAll() (tags []model.Tag, err error) {
	return tags, r.db.GormDB.Order("sort asc").Find(&tags).Error
}

// GetAllIDs get all ids
func (r TagRepository) GetAllStringIDs() (ids []string) {
	r.db.GormDB.Model(&model.Tag{}).Pluck("id", &ids)
	return ids
}

// Save tag
func (r TagRepository) Save(tag model.Tag) (model.Tag, error) {
	return tag, r.db.GormDB.Create(&tag).Error
}

// Update updates tag
func (r TagRepository) Update(tag model.Tag) (model.Tag, error) {
	return tag, r.db.GormDB.Save(&tag).Error
}

// GetOne gets ont tag
func (r TagRepository) GetOne(id uint) (tag model.Tag, err error) {
	return tag, r.db.GormDB.Where("id = ?", id).First(&tag).Error
}

// GetOneByEmail is find by name
func (r TagRepository) GetOneByName(name string) (tag model.Tag, err error) {
	return tag, r.db.GormDB.Where("name = ?", name).First(&tag).Error
}

// Delete deletes the row of data
func (r TagRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Tag{}).Error
}

// MaxSort max sort value
func (r TagRepository) MaxSort() (max uint) {
	r.db.GormDB.Model(&model.Tag{}).Select("max(sort)").Scan(&max)
	return max
}

// ExistsByName is exists by email
func (r TagRepository) ExistsByName(name string) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Tag{}).
		Where("name = ?", name).
		Count(&count).
		Error

	return (count > 0), err
}

// ExistsByIDName is exists by id and name
func (r TagRepository) ExistsByIDName(id uint, name string) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Tag{}).
		Where("id <> ? AND name = ?", id, name).
		Count(&count).
		Error

	return (count > 0), err
}

func (r TagRepository) SortUpdate(ids []uint) error {
	tags, err := r.GetAll()
	if err != nil {
		return err
	}
	return r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		counter := 1
		// id check
		for _, id := range ids {
			// tag id check
			for _, tag := range tags {
				if tag.ID == id {
					tag.Sort = uint(counter)
					if updateErr := tx.Save(&tag).Error; updateErr != nil {
						return updateErr
					}
					counter++
				}
			}
		}
		return nil
	})
}
