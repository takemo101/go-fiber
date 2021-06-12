package object

import (
	"strings"
)

// DiscussionSearchInput search form to service
type DiscussionSearchInput struct {
	keyword string
	page    int
}

func NewDiscussionSearchInput(
	keyword string,
	page int,
) DiscussionSearchInput {
	return DiscussionSearchInput{
		keyword: keyword,
		page:    page,
	}
}

func (o DiscussionSearchInput) GetKeyword() string {
	return strings.TrimSpace(o.keyword)
}

func (o DiscussionSearchInput) GetPage() int {
	if o.page > 0 {
		return o.page
	}
	return 0
}
