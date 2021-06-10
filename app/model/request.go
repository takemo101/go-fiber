package model

import (
	"strconv"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

// RequestStatus for request
type RequestStatus string

const (
	RequestStatusDraft   RequestStatus = "draft"
	RequestStatusApply   RequestStatus = "apply"
	RequestStatusRemand  RequestStatus = "remand"
	RequestStatusRelease RequestStatus = "release"
	RequestStatusPrivate RequestStatus = "private"
)

func (r RequestStatus) String() string {
	return string(r)
}

func (r RequestStatus) Name() string {
	switch r {
	case RequestStatusDraft:
		return "下書き"
	case RequestStatusApply:
		return "公開申請"
	case RequestStatusRemand:
		return "差し戻し"
	case RequestStatusRelease:
		return "公開中"
	}
	return "公開終了"
}

func ToRequestStatusArray() []KeyName {
	return []KeyName{
		{
			Key:  string(RequestStatusDraft),
			Name: RequestStatusDraft.Name(),
		},
		{
			Key:  string(RequestStatusApply),
			Name: RequestStatusApply.Name(),
		},
		{
			Key:  string(RequestStatusRemand),
			Name: RequestStatusRemand.Name(),
		},
		{
			Key:  string(RequestStatusRelease),
			Name: RequestStatusRelease.Name(),
		},
		{
			Key:  string(RequestStatusPrivate),
			Name: RequestStatusPrivate.Name(),
		},
	}
}

// Request is request request
type Request struct {
	gorm.Model
	Title      string        `gorm:"type:varchar(191);not null"`
	Content    string        `gorm:"type:longtext;not null"`
	Status     RequestStatus `gorm:"type:varchar(20);index;not null;default:draft"`
	Tags       []Tag         `gorm:"many2many:request_tags;"`
	CategoryID uint          `gorm:"index"`
	Category   Category      `gorm:"constraint:OnDelete:SET NULL;"`
	UserID     uint          `gorm:"index"`
	User       User          `gorm:"constraint:OnDelete:SET NULL;"`
	Suggests   []Suggest
}

func (m Request) GetCategoryStringID() string {
	return strconv.Itoa(int(m.CategoryID))
}

func (m Request) GetTagStringIDs() []string {
	sIDs := funk.Map(m.Tags, func(tag Tag) string {
		return strconv.Itoa(int(tag.ID))
	})
	return sIDs.([]string)
}

// RequestTag is request 2 tag
type RequestTag struct {
	RequestID uint    `gorm:"primaryKey"`
	Request   Request `gorm:"constraint:OnDelete:CASCADE;"`
	TagID     uint    `gorm:"primaryKey"`
	Tag       Tag     `gorm:"constraint:OnDelete:CASCADE;"`
}
