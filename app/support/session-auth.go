package support

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
)

// SessionAdminAuth is user authentication
type SessionAdminAuth struct {
	admin      *model.Admin
	repository repository.AdminRepository
}

// NewSessionAdminAuth is create auth support
func NewSessionAdminAuth(
	repository repository.AdminRepository,
) *SessionAdminAuth {
	return &SessionAdminAuth{
		repository: repository,
	}
}

// Attempt is user login
func (a *SessionAdminAuth) Attempt(email string, pass string, session *session.Session) bool {
	if ok := a.AttemptSession(session); ok {
		return ok
	}

	admin, err := a.repository.FindByEmail(email)
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

// AttemptSession session auth
func (a *SessionAdminAuth) AttemptSession(session *session.Session) bool {
	id, ok := session.Get("admin_id").(uint)
	if ok {
		admin, err := a.repository.GetOne(id)
		if err == nil {
			a.admin = &admin
			a.Save(admin.ID, session)
			return true
		}
	}
	return false
}

// Save save id to session
func (a *SessionAdminAuth) Save(id uint, session *session.Session) {
	session.Set("admin_id", id)
	session.Save()
}

// Logout admin logout
func (a *SessionAdminAuth) Logout(session *session.Session) {
	a.admin = nil
	session.Set("admin_id", nil)
	session.Save()
}

// Admin admin model
func (a *SessionAdminAuth) Admin() *model.Admin {
	return a.admin
}

// ID admin id
func (a *SessionAdminAuth) ID() uint {
	return a.admin.ID
}

// Check check auth
func (a *SessionAdminAuth) Check() bool {
	return a.admin != nil
}
