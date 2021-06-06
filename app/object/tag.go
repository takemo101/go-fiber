package object

import (
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

// TagInput input form to service
type TagInput struct {
	name string
}

func NewTagInput(
	name string,
) TagInput {
	return TagInput{
		name: name,
	}
}

func (o TagInput) GetName() string {
	return strings.TrimSpace(o.name)
}

// TagSortInput input form to service
type TagSortInput struct {
	ids []string
}

func NewTagSortInput(
	ids []string,
) TagSortInput {
	return TagSortInput{
		ids: ids,
	}
}

func (o TagSortInput) GetIDs() []uint {
	uintIDs := funk.Map(o.ids, func(id string) uint {
		if uID, err := strconv.Atoi(id); err == nil {
			return uint(uID)
		}
		return 0
	})
	return funk.UniqUInt(uintIDs.([]uint))
}
