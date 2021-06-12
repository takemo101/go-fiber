package object

import (
	"strings"
)

// UserInput input form to service
type UserInput struct {
	name  string
	email string
	pass  string
}

func NewUserInput(
	name string,
	email string,
	pass string,
) UserInput {
	return UserInput{
		name:  name,
		email: email,
		pass:  pass,
	}
}

func (o UserInput) GetName() string {
	return strings.TrimSpace(o.name)
}

func (o UserInput) GetEmail() string {
	return strings.TrimSpace(o.email)
}

func (o UserInput) HasPass() bool {
	return len(o.GetPass()) > 0
}

func (o UserInput) GetPass() []byte {
	return []byte(strings.TrimSpace(o.pass))
}

// UserSearchInput search form to service
type UserSearchInput struct {
	keyword string
	page    int
}

func NewUserSearchInput(
	keyword string,
	page int,
) UserSearchInput {
	return UserSearchInput{
		keyword: keyword,
		page:    page,
	}
}

func (o UserSearchInput) GetKeyword() string {
	return strings.TrimSpace(o.keyword)
}

func (o UserSearchInput) GetPage() int {
	if o.page > 0 {
		return o.page
	}
	return 0
}
