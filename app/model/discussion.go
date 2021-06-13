package model

import (
	"gorm.io/gorm"
)

// DiscussionType for suggest
type DiscussionType string

const (
	DiscussionTypeStart     DiscussionType = "start"
	DiscussionTypeMatch     DiscussionType = "match"
	DiscussionTypeSugesster DiscussionType = "sugesster"
	DiscussionTypeRequester DiscussionType = "requester"
	DiscussionTypeCarryOut  DiscussionType = "carryout"
)

func (r DiscussionType) String() string {
	return string(r)
}

func (r DiscussionType) Name() string {
	switch r {
	case DiscussionTypeStart:
		return "提案"
	case DiscussionTypeMatch:
		return "相談開始"
	case DiscussionTypeSugesster:
		return "提案者から"
	case DiscussionTypeRequester:
		return "依頼者から"
	case DiscussionTypeCarryOut:
		return "依頼達成報告"
	}
	return ""
}

func ToDiscussionTypeArray() []KeyName {
	return []KeyName{
		{
			Key:  string(DiscussionTypeStart),
			Name: DiscussionTypeStart.Name(),
		},
		{
			Key:  string(DiscussionTypeStart),
			Name: DiscussionTypeStart.Name(),
		},
		{
			Key:  string(DiscussionTypeSugesster),
			Name: DiscussionTypeSugesster.Name(),
		},
		{
			Key:  string(DiscussionTypeRequester),
			Name: DiscussionTypeRequester.Name(),
		},
	}
}

func (r DiscussionType) IsSuggester() bool {
	return r == DiscussionTypeStart || r == DiscussionTypeSugesster
}

func (r DiscussionType) IsRequester() bool {
	return !r.IsSuggester()
}

// Discussion is suggest discussion
type Discussion struct {
	gorm.Model
	Type       DiscussionType `gorm:"type:varchar(20);index;not null;default:suggest"`
	Message    string         `gorm:"type:text;not null"`
	IsRead     bool           `gorm:"index"`
	SenderID   uint           `gorm:"index"`
	Sender     User           `gorm:"constraint:OnDelete:SET NULL;"`
	ReceiverID uint           `gorm:"index"`
	Receiver   User           `gorm:"constraint:OnDelete:SET NULL;"`
	SuggestID  uint           `gorm:"index;not null"`
	Suggest    Suggest        `gorm:"constraint:OnDelete:CASCADE;"`
}
