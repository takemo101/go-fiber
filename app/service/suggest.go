package service

import (
	"errors"

	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// SuggestService service logic
type SuggestService struct {
	Repository           repository.SuggestRepository
	RequestRepository    repository.RequestRepository
	DiscussionRepository repository.DiscussionRepository
	logger               pkg.Logger
}

// NewSuggestService new service
func NewSuggestService(
	repository repository.SuggestRepository,
	requestRepository repository.RequestRepository,
	discussionRepository repository.DiscussionRepository,
	logger pkg.Logger,
) SuggestService {
	return SuggestService{
		Repository:           repository,
		RequestRepository:    requestRepository,
		DiscussionRepository: discussionRepository,
		logger:               logger,
	}
}

// GetAll suggests
func (s SuggestService) FindAll() ([]model.Suggest, error) {
	return s.Repository.GetAll()
}

// SendStartMessage create suggest start
func (s SuggestService) SendStartMessage(
	requestID uint,
	senderID uint,
	message string,
) (suggest model.Suggest, err error) {

	request, requestErr := s.RequestRepository.GetOne(requestID)
	if requestErr != nil {
		return suggest, requestErr
	}

	if !request.Status.IsRelease() {
		return suggest, errors.New("request has not been release")
	}

	if request.IsClose {
		return suggest, errors.New("request closed")
	}

	receiverID := request.UserID
	if receiverID == senderID {
		return suggest, errors.New("requester can't start message")
	}

	exists, existsErr := s.Repository.ExistsByRequestIDAndSuggesterID(requestID, senderID)
	if existsErr != nil {
		return suggest, existsErr
	}
	if exists {
		return suggest, errors.New("already suggested")
	}

	return s.Repository.SaveWithDiscussion(
		model.Suggest{
			Status:      model.SuggestStatusStart,
			RequestID:   requestID,
			SuggesterID: senderID,
		},
		model.Discussion{
			Type:       model.DiscussionTypeStart,
			Message:    message,
			IsRead:     false,
			SenderID:   senderID,
			ReceiverID: receiverID,
		},
	)
}

// SendMatchMessage create suggest match
func (s SuggestService) SendMatchMessage(
	suggestID uint,
	senderID uint,
	message string,
) (suggest model.Suggest, err error) {

	suggest, suggestErr := s.Repository.GetOne(suggestID)
	if suggestErr != nil {
		return suggest, suggestErr
	}

	if !suggest.Status.IsStart() {
		return suggest, errors.New("not in start state")
	}

	if suggest.IsCloseAll() {
		return suggest, errors.New("suggest closed")
	}

	receiverID := suggest.SuggesterID
	if receiverID == senderID {
		return suggest, errors.New("suggester can't match message")
	}

	suggest.Status = model.SuggestStatusDiscussion
	_, discussionErr := s.DiscussionRepository.SaveWithUpdateSuggest(
		model.Discussion{
			Type:       model.DiscussionTypeMatch,
			Message:    message,
			IsRead:     false,
			SenderID:   senderID,
			ReceiverID: receiverID,
		},
		suggest,
	)
	return suggest, discussionErr
}

// SendMessage create discussion
func (s SuggestService) SendMessage(
	suggestID uint,
	senderID uint,
	message string,
) (suggest model.Suggest, err error) {

	suggest, suggenstErr := s.Repository.GetOne(suggestID)
	if suggenstErr != nil {
		return suggest, suggenstErr
	}

	if suggest.IsCloseAll() {
		return suggest, errors.New("suggest closed")
	}

	// create need data and check
	receiverID, discussionType, paramErr := s.createDiscussionParameter(
		senderID,
		suggest.SuggesterID,
		suggest.Request.UserID,
	)
	if paramErr != nil {
		return suggest, paramErr
	}

	_, discussionErr := s.DiscussionRepository.Save(model.Discussion{
		Type:       discussionType,
		Message:    message,
		IsRead:     true,
		SenderID:   senderID,
		ReceiverID: receiverID,
		SuggestID:  suggestID,
	})

	return suggest, discussionErr
}

// SendCarryOutMessage create suggest match
func (s SuggestService) SendCarryOutMessage(
	suggestID uint,
	senderID uint,
	message string,
) (suggest model.Suggest, err error) {

	suggest, suggestErr := s.Repository.GetOne(suggestID)
	if suggestErr != nil {
		return suggest, suggestErr
	}

	if !suggest.Status.IsDiscussion() {
		return suggest, errors.New("not in discussion state")
	}

	if suggest.IsCloseAll() {
		return suggest, errors.New("suggest closed")
	}

	receiverID := suggest.SuggesterID
	if receiverID == senderID {
		return suggest, errors.New("suggester can't carryout message")
	}

	suggest.Status = model.SuggestStatusCarryOut
	_, discussionErr := s.DiscussionRepository.SaveWithUpdateSuggest(
		model.Discussion{
			Type:       model.DiscussionTypeRequester,
			Message:    message,
			IsRead:     false,
			SenderID:   senderID,
			ReceiverID: receiverID,
		},
		suggest,
	)
	return suggest, discussionErr
}

// Find get suggest
func (s SuggestService) Find(id uint) (suggest model.Suggest, err error) {
	return s.Repository.GetOne(id)
}

// Find get suggest with discussions
func (s SuggestService) FindWithDiscussions(id uint) (suggest model.Suggest, err error) {
	return s.Repository.GetOneWithDiscussions(id)
}

// Delete remove suggest
func (s SuggestService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

func (s SuggestService) createDiscussionParameter(
	senderID uint,
	suggesterID uint,
	requesterID uint,
) (
	receiverID uint,
	discussionType model.DiscussionType,
	err error,
) {
	if senderID == suggesterID {
		discussionType = model.DiscussionTypeSugesster
	} else if senderID == requesterID {
		receiverID = suggesterID
		discussionType = model.DiscussionTypeRequester
	} else {
		err = errors.New("unauthorized sender user id")
	}

	return receiverID, discussionType, err
}
