package model

import (
	"strings"

	"gorm.io/gorm"
)

// MenuProcess for menu
type MenuProcess string

const (
	MenuProcessNone     MenuProcess = "none"
	MenuProcessMatch    MenuProcess = "match"
	MenuProcessCancel   MenuProcess = "cancel"
	MenuProcessComplete MenuProcess = "complete"
)

func (r MenuProcess) String() string {
	return string(r)
}

func (r MenuProcess) Name() string {
	switch r {
	case MenuProcessNone:
		return "依頼中"
	case MenuProcessMatch:
		return "契約中"
	case MenuProcessCancel:
		return "契約破棄"
	}
	return "契約完了"
}

func MenuProcessFromString(process string) MenuProcess {
	switch strings.ToLower(process) {
	case string(MenuProcessNone):
		return MenuProcessNone
	case string(MenuProcessMatch):
		return MenuProcessMatch
	case string(MenuProcessCancel):
		return MenuProcessCancel
	}
	return MenuProcessComplete
}

func ToMenuProcessArray() []KeyName {
	return []KeyName{
		{
			Key:  string(MenuProcessNone),
			Name: MenuProcessNone.Name(),
		},
		{
			Key:  string(MenuProcessMatch),
			Name: MenuProcessMatch.Name(),
		},
		{
			Key:  string(MenuProcessCancel),
			Name: MenuProcessCancel.Name(),
		},
		{
			Key:  string(MenuProcessComplete),
			Name: MenuProcessComplete.Name(),
		},
	}
}

// MenuStatus for menu
type MenuStatus string

const (
	MenuStatusDraft   MenuStatus = "draft"
	MenuStatusApply   MenuStatus = "apply"
	MenuStatusRemand  MenuStatus = "remand"
	MenuStatusRelease MenuStatus = "release"
	MenuStatusPrivate MenuStatus = "private"
)

func (r MenuStatus) String() string {
	return string(r)
}

func (r MenuStatus) Name() string {
	switch r {
	case MenuStatusDraft:
		return "下書き"
	case MenuStatusApply:
		return "公開申請"
	case MenuStatusRemand:
		return "差し戻し"
	case MenuStatusRelease:
		return "公開中"
	}
	return "非公開"
}

func MenuStatusFromString(status string) MenuStatus {
	switch strings.ToLower(status) {
	case string(MenuStatusDraft):
		return MenuStatusDraft
	case string(MenuStatusApply):
		return MenuStatusApply
	case string(MenuStatusRemand):
		return MenuStatusRemand
	case string(MenuStatusRelease):
		return MenuStatusRelease
	}
	return MenuStatusPrivate
}

func ToMenuStatusArray() []KeyName {
	return []KeyName{
		{
			Key:  string(MenuStatusDraft),
			Name: MenuStatusDraft.Name(),
		},
		{
			Key:  string(MenuStatusApply),
			Name: MenuStatusApply.Name(),
		},
		{
			Key:  string(MenuStatusRemand),
			Name: MenuStatusRemand.Name(),
		},
		{
			Key:  string(MenuStatusRelease),
			Name: MenuStatusRelease.Name(),
		},
		{
			Key:  string(MenuStatusPrivate),
			Name: MenuStatusPrivate.Name(),
		},
	}
}

// Menu is request menu
type Menu struct {
	gorm.Model
	Title      string      `gorm:"type:varchar(191);not null"`
	Content    string      `gorm:"type:longtext;not null"`
	Process    MenuProcess `gorm:"type:varchar(20);index;not null;default:none"`
	Status     MenuStatus  `gorm:"type:varchar(20);index;not null;default:draft"`
	Tags       []Tag       `gorm:"many2many:menu_tags;"`
	CategoryID uint        `gorm:"index"`
	Category   Category    `gorm:"constraint:OnDelete:SET NULL;"`
	UserID     uint        `gorm:"index"`
	User       User        `gorm:"constraint:OnDelete:SET NULL;"`
}

// MenuTag is menu 2 tag
type MenuTag struct {
	MenuID uint `gorm:"primary_key"`
	Menu   Menu `gorm:"constraint:OnDelete:CASCADE;"`
	TagID  uint `gorm:"primary_key"`
	Tag    Tag  `gorm:"constraint:OnDelete:CASCADE;"`
}
