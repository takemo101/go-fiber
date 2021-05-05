package support

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
)

// SessionAuth interface
type SessionAuth interface {
	Attempt(string, string, *session.Session) bool
	AttemptSession(*session.Session) bool
	Save(uint, *session.Session)
	Logout(*session.Session)
	ID() uint
	Check() bool
}

// SessionAdminAuth is user authentication
type SessionAdminAuth struct {
	admin      *model.Admin
	repository repository.AdminRepository
}

// NewSessionAuth is create auth support
func NewSessionAdminAuth(
	repository repository.AdminRepository,
) SessionAuth {
	return &SessionAdminAuth{
		repository: repository,
	}
}

// Attempt is user login
func (a *SessionAdminAuth) Attempt(name string, pass string, session *session.Session) bool {
	if ok := a.AttemptSession(session); ok {
		return ok
	}

	admin, err := a.repository.FindByName(name)
	if err != nil {
		return false
	}

	if !helper.ComparePass(admin.Pass, pass) {
		return false
	}

	a.admin = &admin
	a.Save(admin.ID, session)

	return true
}

// AttemptSession is session auth
func (a *SessionAdminAuth) AttemptSession(session *session.Session) bool {
	id, ok := session.Get("admin_id").(uint)
	if ok {
		admin, err := a.repository.GetOne(id)
		if err == nil {
			a.Save(admin.ID, session)
			return true
		}
	}
	return false
}

// Save is save id to session
func (a *SessionAdminAuth) Save(id uint, session *session.Session) {
	session.Set("admin_id", id)
	session.Save()
}

// Logout is user logout
func (a *SessionAdminAuth) Logout(session *session.Session) {
	a.admin = nil
	session.Delete("admin_id")
	session.Save()
}

// User is user model
func (a *SessionAdminAuth) User() *model.Admin {
	return a.admin
}

// ID is user id
func (a *SessionAdminAuth) ID() uint {
	return a.admin.ID
}

// Check is check auth
func (a *SessionAdminAuth) Check() bool {
	return a.admin != nil
}
