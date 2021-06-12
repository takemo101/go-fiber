package object

import (
	"strings"

	"github.com/takemo101/go-fiber/app/model"
)

// RequestInput input form to service
type RequestInput struct {
	title      string
	content    string
	thumbnail  string
	status     string
	tagIDs     []uint
	categoryID uint
}

func NewRequestInput(
	title string,
	content string,
	thumbnail string,
	status string,
	tagIDs []uint,
	categoryID uint,
) RequestInput {
	return RequestInput{
		title:      title,
		content:    content,
		thumbnail:  thumbnail,
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

func (o RequestInput) GetThumbnail() string {
	return o.thumbnail
}

func (o RequestInput) HasThumbnail() bool {
	thumbnail := o.GetThumbnail()
	return thumbnail != ""
}

func (o RequestInput) GetStatus() model.RequestStatus {
	return model.RequestStatus(o.status)
}

func (o RequestInput) GetTagIDs() []uint {
	return o.tagIDs
}

func (o RequestInput) GetCategoryID() uint {
	return o.categoryID
}

// RequestSearchInput search form to service
type RequestSearchInput struct {
	keyword string
	page    int
}

func NewRequestSearchInput(
	keyword string,
	page int,
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
	if o.page > 0 {
		return o.page
	}
	return 0
}
