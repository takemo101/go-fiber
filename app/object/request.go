package object

import (
	"strconv"
	"strings"

	"github.com/takemo101/go-fiber/app/model"
	"github.com/thoas/go-funk"
)

// RequestInput input form to service
type RequestInput struct {
	title      string
	content    string
	status     string
	tagIDs     []string
	categoryID string
}

func NewRequestInput(
	title string,
	content string,
	status string,
	tagIDs []string,
	categoryID string,
) RequestInput {
	return RequestInput{
		title:      title,
		content:    content,
		status:     status,
		tagIDs:     tagIDs,
		categoryID: categoryID,
	}
}

func (o RequestInput) GetTitle() string {
	return strings.TrimSpace(o.title)
}

func (o RequestInput) GetContent() string {
	return o.content
}

func (o RequestInput) GetStatus() model.RequestStatus {
	return model.RequestStatus(o.status)
}

func (o RequestInput) GetTagIDs() []uint {
	uintIDs := funk.Map(o.tagIDs, func(id string) uint {
		if iID, err := strconv.Atoi(id); err == nil {
			return uint(iID)
		}
		return 0
	})
	return funk.UniqUInt(uintIDs.([]uint))
}

func (o RequestInput) GetCategoryID() uint {
	if id, err := strconv.Atoi(o.categoryID); err == nil {
		return uint(id)
	}
	return 0
}

// RequestSearchInput search form to service
type RequestSearchInput struct {
	keyword string
	page    string
}

func NewRequestSearchInput(
	keyword string,
	page string,
) RequestSearchInput {
	return RequestSearchInput{
		keyword: keyword,
		page:    page,
	}
}

func (o RequestSearchInput) GetKeyword() string {
	return strings.TrimSpace(o.keyword)
}

func (o RequestSearchInput) GetPage() int {
	if page, err := strconv.Atoi(o.page); err == nil {
		return page
	}
	return 0
}
