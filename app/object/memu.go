package object

import (
	"strconv"
	"strings"

	"github.com/takemo101/go-fiber/app/model"
	"github.com/thoas/go-funk"
)

// MenuInput input form to service
type MenuInput struct {
	title      string
	content    string
	process    string
	status     string
	tagIDs     []string
	categoryID string
}

func NewMenuInput(
	title string,
	content string,
	process string,
	status string,
	tagIDs []string,
	categoryID string,
) MenuInput {
	return MenuInput{
		title:      title,
		content:    content,
		process:    process,
		status:     status,
		tagIDs:     tagIDs,
		categoryID: categoryID,
	}
}

func (o MenuInput) GetTitle() string {
	return strings.TrimSpace(o.title)
}

func (o MenuInput) GetContent() string {
	return o.content
}

func (o MenuInput) GetProcess() model.MenuProcess {
	return model.MenuProcess(o.process)
}

func (o MenuInput) GetStatus() model.MenuStatus {
	return model.MenuStatus(o.status)
}

func (o MenuInput) GetTagIDs() []uint {
	uintIDs := funk.Map(o.tagIDs, func(id string) uint {
		if iID, err := strconv.Atoi(id); err == nil {
			return uint(iID)
		}
		return 0
	})
	return funk.UniqUInt(uintIDs.([]uint))
}

func (o MenuInput) GetCategoryID() uint {
	if id, err := strconv.Atoi(o.categoryID); err == nil {
		return uint(id)
	}
	return 0
}

// MenuSearchInput search form to service
type MenuSearchInput struct {
	keyword string
	page    string
}

func NewMenuSearchInput(
	keyword string,
	page string,
) MenuSearchInput {
	return MenuSearchInput{
		keyword: keyword,
		page:    page,
	}
}

func (o MenuSearchInput) GetKeyword() string {
	return strings.TrimSpace(o.keyword)
}

func (o MenuSearchInput) GetPage() int {
	if page, err := strconv.Atoi(o.page); err == nil {
		return page
	}
	return 0
}
