package object

import (
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

// CategoryInput input form to service
type CategoryInput struct {
	name     string
	isActive string
}

func NewCategoryInput(
	name string,
	isActive string,
) CategoryInput {
	return CategoryInput{
		name:     name,
		isActive: isActive,
	}
}

func (o CategoryInput) GetName() string {
	return strings.TrimSpace(o.name)
}

func (o CategoryInput) GetIsActive() bool {
	return StringToBool(o.isActive)
}

// CategorySortInput input form to service
type CategorySortInput struct {
	ids []string
}

func NewCategorySortInput(
	ids []string,
) CategorySortInput {
	return CategorySortInput{
		ids: ids,
	}
}

func (o CategorySortInput) GetIDs() []uint {
	uintIDs := funk.Map(o.ids, func(id string) uint {
		if uID, err := strconv.Atoi(id); err == nil {
			return uint(uID)
		}
		return 0
	})
	return funk.UniqUInt(uintIDs.([]uint))
}
