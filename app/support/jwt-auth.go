package support

import (
	"errors"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/pkg"
)

// JWTUserAuth is user authentication
type JWTUserAuth struct {
	user       *model.User
	config     pkg.Config
	repository repository.UserRepository
}

// NewJWTUserAuth is create auth support
func NewJWTUserAuth(
	repository repository.UserRepository,
	config pkg.Config,
) *JWTUserAuth {
	return &JWTUserAuth{
		repository: repository,
		config:     config,
	}
}

// Attempt is user login
func (a *JWTUserAuth) Attempt(email string, pass string, c *fiber.Ctx) bool {
	if ok := a.AttemptToken(c); ok {
		return true
	}

	user, err := a.repository.GetOneByEmail(email)
	if err != nil {
		return false
	}

	if !helper.ComparePass(user.Pass, pass) {
		return false
	}

	a.user = &user
	return true
}

// AttemptToken jwt auth
func (a *JWTUserAuth) AttemptToken(c *fiber.Ctx) bool {
	claims, claimsErr := a.ExtractClaims(c)
	if claimsErr != nil {
		return false
	}

	tempID, keyOK := claims["id"]

	if keyOK {
		if id, ok := tempID.(float64); ok {
			user, err := a.repository.GetOne(uint(id))
			if err == nil {
				a.user = &user
				return true
			}
		}
	}

	return false
}

// ExtractClaims context to claims
func (a *JWTUserAuth) ExtractClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	if user, castOK := c.Locals(a.config.JWT.Context.Key).(*jwt.Token); castOK {
		if claims, ok := user.Claims.(jwt.MapClaims); ok {
			return claims, nil
		}
	}

	return nil, errors.New("cast failed")
}

// GenerateTokenByUser create jwt token by user model
func (a *JWTUserAuth) GenerateTokenByUser(user *model.User) (string, error) {
	// new token
	token := jwt.New(jwt.GetSigningMethod(a.config.JWT.Signing.Method))

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * a.config.JWT.Expiration).Unix()

	// generate encoded token
	t, tokenErr := token.SignedString([]byte(a.config.JWT.Signing.Key))
	if tokenErr != nil {
		return t, tokenErr
	}

	return t, nil
}

// User user model
func (a *JWTUserAuth) User() *model.User {
	return a.user
}

// ID user id
func (a *JWTUserAuth) ID() uint {
	if a.Check() {
		return a.user.ID
	}
	return 0
}

// Check check auth
func (a *JWTUserAuth) Check() bool {
	return a.user != nil
}
