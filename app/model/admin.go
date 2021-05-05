package model

import (
	"github.com/takemo101/go-fiber/app/helper"
	"gorm.io/gorm"
)

// User is auth user
type Admin struct {
	gorm.Model
	Name  string
	Email string
	Pass  []byte
}

// BeforeSave is Hook
func (u *Admin) BeforeSave(db *gorm.DB) (err error) {
	if len(u.Pass) > 0 {
		hash, err := helper.CreatePass(string(u.Pass))
		if err != nil {
			return err
		}
		u.Pass = []byte(hash)
	}
	return
}
