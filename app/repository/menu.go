package repository

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
	"gorm.io/gorm"
)

// MenuRepository database structure
type MenuRepository struct {
	db     pkg.Database
	logger pkg.Logger
}

// NewMenuRepository creates a new menu repository
func NewMenuRepository(db pkg.Database, logger pkg.Logger) MenuRepository {
	return MenuRepository{
		db:     db,
		logger: logger,
	}
}

// GetAll gets all menus
func (r MenuRepository) GetAll() (menus []model.Menu, err error) {
	return menus, r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User").
		Find(&menus).
		Error
}

// Save menu
func (r MenuRepository) Save(menu model.Menu) (model.Menu, error) {
	return menu, r.db.GormDB.Create(&menu).Error
}

// SaveWithTagIDs menu and menu_tag
func (r MenuRepository) SaveWithTagIDs(menu model.Menu, tagIDs []uint) (model.Menu, error) {
	if err := r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		if menuErr := tx.Create(&menu).Error; menuErr != nil {
			return menuErr
		}

		var menuTags = make([]model.MenuTag, len(tagIDs))
		for index, id := range tagIDs {
			menuTags[index] = model.MenuTag{
				MenuID: menu.ID,
				TagID:  id,
			}
		}
		return tx.Create(&menuTags).Error
	}); err != nil {
		return menu, err
	}
	return menu, nil
}

// Update updates menu
func (r MenuRepository) Update(menu model.Menu) (model.Menu, error) {
	return menu, r.db.GormDB.Save(&menu).Error
}

// UpdateWithTagIDs menu and menu_tag
func (r MenuRepository) UpdateWithTagIDs(menu model.Menu, tagIDs []uint) (model.Menu, error) {
	if err := r.db.GormDB.Transaction(func(tx *gorm.DB) error {
		if deleteErr := tx.Where("menu_id = ?", menu.ID).Delete(&model.MenuTag{}).Error; deleteErr != nil {
			return deleteErr
		}

		if menuErr := tx.Create(&menu).Error; menuErr != nil {
			return menuErr
		}

		var menuTags = make([]model.MenuTag, len(tagIDs))
		for index, id := range tagIDs {
			menuTags[index] = model.MenuTag{
				MenuID: menu.ID,
				TagID:  id,
			}
		}
		return tx.Create(&menuTags).Error
	}); err != nil {
		return menu, err
	}
	return menu, nil
}

// GetOne gets ont menu
func (r MenuRepository) GetOne(id uint) (menu model.Menu, err error) {
	return menu, r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User").
		Where("id = ?", id).
		First(&menu).
		Error
}

// Delete deletes the row of data
func (r MenuRepository) Delete(id uint) error {
	return r.db.GormDB.Where("id = ?", id).Delete(&model.Menu{}).Error
}
