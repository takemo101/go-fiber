package service

import (
	"errors"
	"fmt"

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

// SendSuggestMessage create suggest
func (s SuggestService) SendSuggestMessage(
	requestID uint,
	senderID uint,
	message string,
) (suggest model.Suggest, err error) {

	request, requestErr := s.RequestRepository.GetOne(requestID)
	if requestErr != nil {
		return suggest, requestErr
	}

	receiverID := request.UserID

	if receiverID == senderID {
		return suggest, errors.New("requester can't suggest")
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
			Status:      model.SuggestStatusDiscussion,
			RequestID:   requestID,
			SuggesterID: senderID,
		},
		model.Discussion{
			Type:       model.DiscussionTypeSuggest,
			Message:    message,
			IsRead:     false,
			SenderID:   senderID,
			ReceiverID: receiverID,
		},
	)
}

// ReplyDecline reply decline discussion
func (s SuggestService) ReplyDecline(id uint, replyerID uint, accept bool) (model.Suggest, error) {
	suggest, err := s.Find(id)
	if err != nil {
		return suggest, err
	}

	if suggest.Status == model.SuggestStatusEnd {
		return suggest, errors.New("suggest has already ended")
	} else if suggest.Status == model.SuggestStatusRequesterDecline {
		if replyerID != suggest.SuggesterID {
			return suggest, errors.New("not a suggester can't decline reply")
		}
	} else if suggest.Status == model.SuggestStatusSuggesterDecline {
		fmt.Println(suggest.Request.UserID)
		if replyerID != suggest.Request.UserID {
			return suggest, errors.New("not a requrester can't decline reply")
		}
	} else {
		return suggest, errors.New("not ready to decline reply")
	}

	if accept {
		suggest.Status = model.SuggestStatusEnd
	} else {
		suggest.Status = model.SuggestStatusDiscussion
	}

	return s.Repository.Update(suggest)
}

// Find get suggest
func (s SuggestService) Find(id uint) (suggest model.Suggest, err error) {
	return s.Repository.GetOne(id)
}

// Delete remove suggest
func (s SuggestService) Delete(id uint) error {
	return s.Repository.Delete(id)
}
