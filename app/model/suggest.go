package model

import (
	"gorm.io/gorm"
)

// SuggestStatus for suggest
type SuggestStatus string

const (
	SuggestStatusDiscussion       SuggestStatus = "discussion"
	SuggestStatusRequesterDecline SuggestStatus = "decline:requester"
	SuggestStatusSuggesterDecline SuggestStatus = "decline:suggester"
	SuggestStatusEnd              SuggestStatus = "end"
)

func (r SuggestStatus) String() string {
	return string(r)
}

func (r SuggestStatus) Name() string {
	switch r {
	case SuggestStatusDiscussion:
		return "提案中"
	case SuggestStatusRequesterDecline:
		return "終了リクエスト：募集者"
	case SuggestStatusSuggesterDecline:
		return "終了リクエスト：提案者"
	}
	return "提案終了"
}

func (r SuggestStatus) IsAlreadyDeclined() bool {
	return r == SuggestStatusRequesterDecline || r == SuggestStatusSuggesterDecline
}

func ToSuggestStatusArray() []KeyName {
	return []KeyName{
		{
			Key:  string(SuggestStatusDiscussion),
			Name: SuggestStatusDiscussion.Name(),
		},
		{
			Key:  string(SuggestStatusRequesterDecline),
			Name: SuggestStatusRequesterDecline.Name(),
		},
		{
			Key:  string(SuggestStatusSuggesterDecline),
			Name: SuggestStatusSuggesterDecline.Name(),
		},
		{
			Key:  string(SuggestStatusEnd),
			Name: SuggestStatusEnd.Name(),
		},
	}
}

// Suggest is request suggest
type Suggest struct {
	gorm.Model
	Status      SuggestStatus `gorm:"type:varchar(20);index;not null;default:discussion"`
	RequestID   uint          `gorm:"index;not null"`
	Request     Request       `gorm:"constraint:OnDelete:CASCADE;"`
	SuggesterID uint          `gorm:"index"`
	Suggester   User          `gorm:"constraint:OnDelete:SET NULL;"`
}
