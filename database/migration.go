package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations = []*gormigrate.Migration{
	// create admins
	{
		ID: "202106050001",
		Migrate: func(tx *gorm.DB) error {
			type Admin struct {
				gorm.Model
				Name  string `gorm:"type:varchar(191);index;not null"`
				Email string `gorm:"type:varchar(191);uniqueIndex;not null"`
				Pass  []byte
				Role  string `gorm:"type:varchar(191);index;not null;default:admin"`
			}
			return tx.AutoMigrate(&Admin{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Admin struct{}
			return tx.Migrator().DropTable(&Admin{})
		},
	},
	// create users
	{
		ID: "202106050002",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				gorm.Model
				Name  string `gorm:"type:varchar(191);index;not null"`
				Email string `gorm:"type:varchar(191);uniqueIndex;not null"`
				Pass  []byte
			}
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			type User struct{}
			return tx.Migrator().DropTable(&User{})
		},
	},
	// create todos
	{
		ID: "202106050003",
		Migrate: func(tx *gorm.DB) error {
			type Admin struct {
				gorm.Model
			}
			type Todo struct {
				gorm.Model
				Text    string `gorm:"type:text;not null"`
				Status  string `gorm:"type:varchar(191);index;not null;default:up"`
				AdminID uint   `gorm:"index"`
				Admin   Admin  `gorm:"constraint:OnDelete:CASCADE;"`
			}
			return tx.AutoMigrate(&Todo{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Todo struct{}
			return tx.Migrator().DropTable(&Todo{})
		},
	},
	// create categories
	{
		ID: "202106050004",
		Migrate: func(tx *gorm.DB) error {
			type Category struct {
				gorm.Model
				Name     string `gorm:"type:varchar(191);uniqueIndex;not null"`
				Sort     uint   `gorm:"uniqueIndex;default:1"`
				IsActive bool   `gorm:"index;default:true"`
			}
			return tx.AutoMigrate(&Category{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Category struct{}
			return tx.Migrator().DropTable(&Category{})
		},
	},
	// create tags
	{
		ID: "202106050005",
		Migrate: func(tx *gorm.DB) error {
			type Tag struct {
				gorm.Model
				Name string `gorm:"type:varchar(191);uniqueIndex;not null"`
				Sort uint   `gorm:"uniqueIndex;default:1"`
			}
			return tx.AutoMigrate(&Tag{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Tag struct{}
			return tx.Migrator().DropTable(&Tag{})
		},
	},
	// create menus
	{
		ID: "202106050006",
		Migrate: func(tx *gorm.DB) error {
			type Category struct {
				gorm.Model
			}
			type User struct {
				gorm.Model
			}
			type Menu struct {
				gorm.Model
				Title      string   `gorm:"type:varchar(191);not null"`
				Content    string   `gorm:"type:longtext;not null"`
				Process    string   `gorm:"type:varchar(20);index;not null;default:none"`
				Status     string   `gorm:"type:varchar(20);index;not null;default:draft"`
				CategoryID uint     `gorm:"index"`
				Category   Category `gorm:"constraint:OnDelete:SET NULL;"`
				UserID     uint     `gorm:"index"`
				User       User     `gorm:"constraint:OnDelete:SET NULL;"`
			}
			return tx.AutoMigrate(&Menu{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Menu struct{}
			return tx.Migrator().DropTable(&Menu{})
		},
	},
	// create menu_tags
	{
		ID: "202106050007",
		Migrate: func(tx *gorm.DB) error {
			type Menu struct {
				gorm.Model
			}
			type Tag struct {
				gorm.Model
			}
			type MenuTag struct {
				MenuID uint `gorm:"primaryKey"`
				Menu   Menu `gorm:"constraint:OnDelete:CASCADE;"`
				TagID  uint `gorm:"primaryKey"`
				Tag    Tag  `gorm:"constraint:OnDelete:CASCADE;"`
			}
			return tx.AutoMigrate(&MenuTag{})
		},
		Rollback: func(tx *gorm.DB) error {
			type MenuTag struct{}
			return tx.Migrator().DropTable(&MenuTag{})
		},
	},
}
