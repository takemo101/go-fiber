package model

import (
	"strings"

	"gorm.io/gorm"
)

// Role for admin
type Role string

const (
	RoleSystem Role = "system"
	RoleAdmin  Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

func (r Role) Name() string {
	switch r {
	case RoleSystem:
		return "システム管理者"
	case RoleAdmin:
		return "通常管理者"
	}
	return ""
}

func RoleFromString(role string) Role {
	switch strings.ToLower(role) {
	case string(RoleSystem):
		return RoleSystem
	}
	return RoleAdmin
}

func ToRoleArray() []KeyName {
	return []KeyName{
		{
			Key:  RoleSystem,
			Name: RoleSystem.Name(),
		},
		{
			Key:  RoleAdmin,
			Name: RoleAdmin.Name(),
		},
	}
}

// User is auth user
type Admin struct {
	gorm.Model
	Name  string `gorm:"type:varchar(191);index;not null"`
	Email string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Pass  []byte
	Role  Role `gorm:"type:varchar(191);index;not null;default:admin"`
	Todos []Todo
}
