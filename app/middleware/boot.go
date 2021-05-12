package middleware

import (
	"github.com/takemo101/go-fiber/pkg/contract"
	"go.uber.org/fx"
)

// Module export
var Module = fx.Options(
	fx.Provide(NewSession),
	fx.Provide(NewSecure),
	fx.Provide(NewCsrf),
	fx.Provide(NewCors),
	fx.Provide(NewSessionAdminAuth),
	fx.Provide(NewMiddleware),
)

// Middlewares is slice
type Middlewares []contract.Middleware

// NewMiddleware is setup new middlewares
func NewMiddleware(
	session Session,
	secure Secure,

) Middlewares {
	return Middlewares{
		session,
		secure,
	}
}

// Setup all the middleware
func (middlewares Middlewares) Setup() {
	for _, middleware := range middlewares {
		middleware.Setup()
	}
}
