package object

import (
	"strings"

	"github.com/takemo101/go-fiber/app/model"
)

// AdminInput input form to service
type AdminInput struct {
	name  string
	email string
	pass  string
	role  string
}

func NewAdminInput(
	name string,
	email string,
	pass string,
	role string,
) AdminInput {
	return AdminInput{
		name:  name,
		email: email,
		pass:  pass,
		role:  role,
	}
}

func (o AdminInput) GetName() string {
	return strings.TrimSpace(o.name)
}

func (o AdminInput) GetEmail() string {
	return strings.TrimSpace(o.email)
}

func (o AdminInput) HasPass() bool {
	return len(o.GetPass()) > 0
}

func (o AdminInput) GetPass() []byte {
	return []byte(strings.TrimSpace(o.pass))
}

func (o AdminInput) HasRole() bool {
	return len(o.role) > 0
}

func (o AdminInput) GetRole() model.Role {
	return model.Role(o.role)
}

// AdminSearchInput search form to service
type AdminSearchInput struct {
	keyword string
	page    int
}

func NewAdminSearchInput(
	keyword string,
	page int,
) AdminSearchInput {
	return AdminSearchInput{
		keyword: keyword,
		page:    page,
	}
}

func (o AdminSearchInput) GetKeyword() string {
	return strings.TrimSpace(o.keyword)
}

func (o AdminSearchInput) GetPage() int {
	if o.page > 0 {
		return o.page
	}
	return 0
}
