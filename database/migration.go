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
				Sort     uint   `gorm:"index;default:1"`
				IsActive bool   `gorm:"index"`
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
				Sort uint   `gorm:"index;default:1"`
			}
			return tx.AutoMigrate(&Tag{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Tag struct{}
			return tx.Migrator().DropTable(&Tag{})
		},
	},
	// create requests
	{
		ID: "202106050006",
		Migrate: func(tx *gorm.DB) error {
			type Category struct {
				gorm.Model
			}
			type User struct {
				gorm.Model
			}
			type Request struct {
				gorm.Model
				Title      string   `gorm:"type:varchar(191);not null"`
				Content    string   `gorm:"type:longtext;not null"`
				Status     string   `gorm:"type:varchar(20);index;not null;default:draft"`
				CategoryID uint     `gorm:"index"`
				Category   Category `gorm:"constraint:OnDelete:SET NULL;"`
				UserID     uint     `gorm:"index"`
				User       User     `gorm:"constraint:OnDelete:SET NULL;"`
			}
			return tx.AutoMigrate(&Request{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Request struct{}
			return tx.Migrator().DropTable(&Request{})
		},
	},
	// create request_tags
	{
		ID: "202106050007",
		Migrate: func(tx *gorm.DB) error {
			type Request struct {
				gorm.Model
			}
			type Tag struct {
				gorm.Model
			}
			type RequestTag struct {
				RequestID uint    `gorm:"primaryKey"`
				Request   Request `gorm:"constraint:OnDelete:CASCADE;"`
				TagID     uint    `gorm:"primaryKey"`
				Tag       Tag     `gorm:"constraint:OnDelete:CASCADE;"`
			}
			return tx.AutoMigrate(&RequestTag{})
		},
		Rollback: func(tx *gorm.DB) error {
			type RequestTag struct{}
			return tx.Migrator().DropTable(&RequestTag{})
		},
	},
	// create suggest
	{
		ID: "202106050008",
		Migrate: func(tx *gorm.DB) error {
			type Request struct {
				gorm.Model
			}
			type User struct {
				gorm.Model
			}
			type Suggest struct {
				gorm.Model
				Status      string  `gorm:"type:varchar(20);index;not null;default:discussion"`
				IsClose     bool    `gorm:"index;default:false"`
				RequestID   uint    `gorm:"index;not null"`
				Request     Request `gorm:"constraint:OnDelete:CASCADE;"`
				SuggesterID uint    `gorm:"index"`
				Suggester   User    `gorm:"constraint:OnDelete:SET NULL;"`
			}

			return tx.AutoMigrate(&Suggest{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Suggest struct{}
			return tx.Migrator().DropTable(&Suggest{})
		},
	},
	// create discussion
	{
		ID: "202106050009",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				gorm.Model
			}
			type Suggest struct {
				gorm.Model
			}
			type Discussion struct {
				gorm.Model
				Type       string  `gorm:"type:varchar(20);index;not null;default:suggest"`
				Message    string  `gorm:"type:text;not null"`
				IsRead     bool    `gorm:"index"`
				SenderID   uint    `gorm:"index"`
				Sender     User    `gorm:"constraint:OnDelete:SET NULL;"`
				ReceiverID uint    `gorm:"index"`
				Receiver   User    `gorm:"constraint:OnDelete:SET NULL;"`
				SuggestID  uint    `gorm:"index;not null"`
				Suggest    Suggest `gorm:"constraint:OnDelete:CASCADE;"`
			}

			return tx.AutoMigrate(&Discussion{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Discussion struct{}
			return tx.Migrator().DropTable(&Discussion{})
		},
	},
	// add column request
	{
		ID: "202106050010",
		Migrate: func(tx *gorm.DB) error {
			type Request struct {
				IsClose   bool `gorm:"index;default:false"`
				Thumbnail string
			}
			return tx.AutoMigrate(&Request{})
		},
		Rollback: func(tx *gorm.DB) error {
			for _, column := range []string{
				"thumbnail",
				"is_close",
			} {
				migrator := tx.Migrator()
				if err := migrator.DropColumn("requests", column); err != nil {
					return err
				}
			}
			return nil
		},
	},
	// add column user
	{
		ID: "202106050011",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				CarryOutCounter string `gorm:"index;default:0"`
			}
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn("users", "carryout_counter")
		},
	},
}
