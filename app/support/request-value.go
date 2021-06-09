package support

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
)

const (
	JWTUserAuthKey      string = "jwt-user-auth"
	SessionAdminAuthKey string = "session-admin-auth"
	ViewRenderKey       string = "view-render"
	ResponseHelperKey   string = "response-helper"
)

// RequestValue
type RequestValue struct {
	//
}

// NewRequestValue
func NewRequestValue() RequestValue {
	return RequestValue{}
}

// GetJWTUserAuth
func (r RequestValue) GetJWTUserAuth(c *fiber.Ctx) *JWTUserAuth {
	if auth, ok := c.Locals(JWTUserAuthKey).(*JWTUserAuth); ok {
		return auth
	}
	return nil
}

// SetJWTUserAuth
func (r RequestValue) SetJWTUserAuth(c *fiber.Ctx, auth *JWTUserAuth) {
	c.Locals(JWTUserAuthKey, auth)
}

// GetSessionAdminAuth
func (r RequestValue) GetSessionAdminAuth(c *fiber.Ctx) *SessionAdminAuth {
	if auth, ok := c.Locals(SessionAdminAuthKey).(*SessionAdminAuth); ok {
		return auth
	}
	return nil
}

// SetJWTUserAuth
func (r RequestValue) SetSessionAdminAuth(c *fiber.Ctx, auth *SessionAdminAuth) {
	c.Locals(SessionAdminAuthKey, auth)
}

// GetViewRender
func (r RequestValue) GetViewRender(c *fiber.Ctx) *helper.ViewRender {
	if render, ok := c.Locals(ViewRenderKey).(*helper.ViewRender); ok {
		return render
	}
	return nil
}

// SetViewRender
func (r RequestValue) SetViewRender(c *fiber.Ctx, render *helper.ViewRender) {
	c.Locals(ViewRenderKey, render)
}

// GetResponseHelper
func (r RequestValue) GetResponseHelper(c *fiber.Ctx) *helper.ResponseHelper {
	if response, ok := c.Locals(ResponseHelperKey).(*helper.ResponseHelper); ok {
		return response
	}
	return nil
}

// SetResponseHelper
func (r RequestValue) SetResponseHelper(c *fiber.Ctx, response *helper.ResponseHelper) {
	c.Locals(ResponseHelperKey, response)
}
