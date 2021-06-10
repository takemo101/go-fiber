package model

import (
	"gorm.io/gorm"
)

// DiscussionType for suggest
type DiscussionType string

const (
	DiscussionTypeSuggest          DiscussionType = "suggest"
	DiscussionTypeSugessterMessage DiscussionType = "sugesster_message"
	DiscussionTypeRequesterMessage DiscussionType = "requester_message"
	DiscussionTypeDecline          DiscussionType = "decline"
)

func (r DiscussionType) String() string {
	return string(r)
}

func (r DiscussionType) Name() string {
	switch r {
	case DiscussionTypeSuggest:
		return "提案"
	case DiscussionTypeSugessterMessage:
		return "提案者メッセージ"
	case DiscussionTypeRequesterMessage:
		return "募集者メッセージ"
	}
	return "終了リクエスト"
}

func ToDiscussionTypeArray() []KeyName {
	return []KeyName{
		{
			Key:  string(DiscussionTypeSuggest),
			Name: DiscussionTypeSuggest.Name(),
		},
		{
			Key:  string(DiscussionTypeSugessterMessage),
			Name: DiscussionTypeSugessterMessage.Name(),
		},
		{
			Key:  string(DiscussionTypeRequesterMessage),
			Name: DiscussionTypeRequesterMessage.Name(),
		},
		{
			Key:  string(DiscussionTypeDecline),
			Name: DiscussionTypeDecline.Name(),
		},
	}
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
