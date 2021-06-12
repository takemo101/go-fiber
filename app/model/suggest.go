package model

import (
	"gorm.io/gorm"
)

// SuggestStatus for suggest
type SuggestStatus string

const (
	SuggestStatusStart      SuggestStatus = "start"
	SuggestStatusDiscussion SuggestStatus = "discussion"
	SuggestStatusCarryOut   SuggestStatus = "carryout"
)

func (r SuggestStatus) String() string {
	return string(r)
}

func (r SuggestStatus) Name() string {
	switch r {
	case SuggestStatusStart:
		return "提案中"
	case SuggestStatusDiscussion:
		return "相談中"
	}
	return "達成"
}

func ToSuggestStatusArray() []KeyName {
	return []KeyName{
		{
			Key:  string(SuggestStatusStart),
			Name: SuggestStatusStart.Name(),
		},
		{
			Key:  string(SuggestStatusDiscussion),
			Name: SuggestStatusDiscussion.Name(),
		},
		{
			Key:  string(SuggestStatusCarryOut),
			Name: SuggestStatusCarryOut.Name(),
		},
	}
}

func (r SuggestStatus) IsStart() bool {
	return r == SuggestStatusStart
}

func (r SuggestStatus) IsDiscussion() bool {
	return r == SuggestStatusDiscussion
}

// Suggest is request suggest
type Suggest struct {
	gorm.Model
	Status      SuggestStatus `gorm:"type:varchar(20);index;not null;default:discussion"`
	IsClose     bool          `gorm:"index;default:false"`
	RequestID   uint          `gorm:"index;not null"`
	Request     Request       `gorm:"constraint:OnDelete:CASCADE;"`
	SuggesterID uint          `gorm:"index"`
	Suggester   User          `gorm:"constraint:OnDelete:SET NULL;"`
	Discussions []Discussion
}

func (m Suggest) IsCloseAll() bool {
	return m.IsClose || m.Request.IsClose
}
