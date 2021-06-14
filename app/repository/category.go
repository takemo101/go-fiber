package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
	"gorm.io/gorm"
)

// CategoryRepository database structure
type CategoryRepository struct {
	db pkg.Database
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db pkg.Database) CategoryRepository {
	return CategoryRepository{
		db: db,
	}
}

// GetAll get all categorys
func (r CategoryRepository) GetAll() (categorys []model.Category, err error) {
	return categorys, r.db.GormDB.Order("sort asc").Find(&categorys).Error
}

// GetAllIDs get all ids
func (r CategoryRepository) GetAllStringIDs() (ids []string) {
	r.db.GormDB.Model(&model.Category{}).Pluck("id", &ids)
	return ids
}

// Save category
func (r CategoryRepository) Save(category model.Category) (model.Category, error) {
	return category, r.db.GormDB.Create(&category).Error
}

// Update updates category
func (r CategoryRepository) Update(category model.Category) (model.Category, error) {
	return category, r.db.GormDB.Save(&category).Error
}

// GetOne get one category
func (r CategoryRepository) GetOne(id uint) (category model.Category, err error) {
	return category, r.db.GormDB.Where("id = ?", id).First(&category).Error
}

// GetOneByEmail is find by name
func (r CategoryRepository) GetOneByName(name string) (category model.Category, err error) {
	return category, r.db.GormDB.Where("name = ?", name).First(&category).Error
}

// Delete deletes the row of data
func (r CategoryRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Category{}).Error
}

// MaxSort max sort value
func (r CategoryRepository) MaxSort() uint {
	var max interface{}
	r.db.GormDB.Model(&model.Category{}).Select("max(sort)").Scan(&max)
	if max == nil {
		return 0
	}
	return max.(uint)
}

// ExistsByName is exists by email
func (r CategoryRepository) ExistsByName(name string) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Category{}).
		Where("name = ?", name).
		Count(&count).
		Error

	return (count > 0), err
}

// ExistsByIDName is exists by id and name
func (r CategoryRepository) ExistsByIDName(id uint, name string) (bool, error) {
	count := int64(0)
	err := r.db.GormDB.Model(&model.Category{}).
		Where("id <> ? AND name = ?", id, name).
		Count(&count).
		Error

	return (count > 0), err
}

func (r CategoryRepository) SortUpdate(ids []uint) error {
	categories, err := r.GetAll()
	if err != nil {
		return err
	}
	return r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		counter := 1
		// id check
		for _, id := range ids {
			// category id check
			for _, category := range categories {
				if category.ID == id {
					category.Sort = uint(counter)
					if updateErr := tx.Save(&category).Error; updateErr != nil {
						return updateErr
					}
					counter++
				}
			}
		}
		return nil
	})
}
