package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is auth user
type Admin struct {
	gorm.Model
	Name string
	Pass []byte
}

// BeforeSave is Hook
func (u *Admin) BeforeSave(db *gorm.DB) (err error) {
	if len(u.Pass) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Pass = hash
	}
	return
}
