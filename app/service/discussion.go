package service

import (
	"errors"

	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// DiscussionService service logic
type DiscussionService struct {
	Repository        repository.DiscussionRepository
	SuggestRepository repository.SuggestRepository
	logger            pkg.Logger
}

// NewDiscussionService new service
func NewDiscussionService(
	repository repository.DiscussionRepository,
	suggestRepository repository.SuggestRepository,
	logger pkg.Logger,
) DiscussionService {
	return DiscussionService{
		Repository:        repository,
		SuggestRepository: suggestRepository,
		logger:            logger,
	}
}

// GetAll discussions
func (s DiscussionService) FindAll() ([]model.Discussion, error) {
	return s.Repository.GetAll()
}

// SendMessage create discussion
func (s DiscussionService) SendMessage(
	suggestID uint,
	senderID uint,
	message string,
) (discussion model.Discussion, err error) {

	suggest, suggenstErr := s.SuggestRepository.GetOne(suggestID)
	if suggenstErr != nil {
		return discussion, suggenstErr
	}

	if statusErr := s.CheckSuggestStatus(suggest.Status); statusErr != nil {
		return discussion, statusErr
	}

	// create need data and check
	receiverID, discussionType, _, dataErr := s.CreateDiscussionData(
		senderID,
		suggest.SuggesterID,
		suggest.Request.UserID,
	)
	if dataErr != nil {
		return discussion, dataErr
	}

	return s.Repository.Save(model.Discussion{
		Type:       discussionType,
		Message:    message,
		IsRead:     true,
		SenderID:   senderID,
		ReceiverID: receiverID,
		SuggestID:  suggestID,
	})
}

// SendMessage create discussion
func (s DiscussionService) SendDeclineMessage(
	suggestID uint,
	senderID uint,
	message string,
) (discussion model.Discussion, err error) {

	suggest, suggenstErr := s.SuggestRepository.GetOne(suggestID)
	if suggenstErr != nil {
		return discussion, suggenstErr
	}

	if statusErr := s.CheckSuggestStatus(suggest.Status); statusErr != nil {
		return discussion, statusErr
	}
	if declineErr := s.CheckSuggestDeclineStatus(suggest.Status); declineErr != nil {
		return discussion, declineErr
	}

	// create need data and check
	receiverID, discussionType, suggestStatus, dataErr := s.CreateDiscussionData(
		senderID,
		suggest.SuggesterID,
		suggest.Request.UserID,
	)
	if dataErr != nil {
		return discussion, dataErr
	}

	suggest.Status = suggestStatus

	discussion, err = s.Repository.SaveWithUpdateSuggest(model.Discussion{
		Type:       discussionType,
		Message:    message,
		IsRead:     true,
		SenderID:   senderID,
		ReceiverID: receiverID,
		SuggestID:  suggestID,
	}, suggest)
	return discussion, err
}

// Find get discussion
func (s DiscussionService) Find(id uint) (discussion model.Discussion, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove discussion
func (s DiscussionService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

// CheckSuggestStatus check suggest status end
func (s DiscussionService) CheckSuggestStatus(status model.SuggestStatus) error {
	if status == model.SuggestStatusEnd {
		return errors.New("suggest has already ended")
	}
	return nil
}

// CheckSuggestStatus check suggest status decline
func (s DiscussionService) CheckSuggestDeclineStatus(status model.SuggestStatus) error {
	if status.IsAlreadyDeclined() {
		return errors.New("suggest has already declined")
	}
	return nil
}

func (s DiscussionService) CreateDiscussionData(
	senderID uint,
	suggesterID uint,
	requesterID uint,
) (
	receiverID uint,
	discussionType model.DiscussionType,
	suggestStatus model.SuggestStatus,
	err error,
) {
	if senderID == suggesterID {
		receiverID = requesterID
		discussionType = model.DiscussionTypeSugessterMessage
		suggestStatus = model.SuggestStatusSuggesterDecline
	} else if senderID == requesterID {
		receiverID = suggesterID
		discussionType = model.DiscussionTypeRequesterMessage
		suggestStatus = model.SuggestStatusRequesterDecline
	} else {
		err = errors.New("unauthorized sender user id")
	}

	return receiverID, discussionType, suggestStatus, err
}
