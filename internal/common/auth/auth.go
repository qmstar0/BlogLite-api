package auth

import (
	"blog/pkg/jwtAuth"
	"common/e"
	"common/server/httperr"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

type User struct {
	Name  string
	Email string
	ID    string
	Role  int16
}

const (
	AuthHeaderKey = "Authorization"
	BearerScopes  = "bearer.Scopes"
	Audience      = "User"
	Issuer        = "Authorization"
	Subject       = "UserToken"
)

// VerificationFunc 该函数做首次验证
// 当 return true, nil 时，表示验证成功，无需继续验证
// 当 return false, nil 时， 表示需要继续验证
// 当 return xxx, err 时，将直接返回错误
type VerificationFunc func(auth string, ctx context.Context) (bool, error)

func tokenFromHeader(r *http.Request) string {
	return r.Header.Get(AuthHeaderKey)
}

const userContextKey = "__userContextKey"

var NoUserInContextError = errors.New("no user in context")

func GetUserFromCtx(c echo.Context) (*User, error) {
	u, ok := c.Get(userContextKey).(*User)
	if ok {
		return u, nil
	}

	return nil, NoUserInContextError
}
func AuthMiddleware() {
	privateFilepath := os.Getenv("AUTH_PRIVATE_FILEPATH")
	if privateFilepath == "" {
		panic("auth middleware is not configured: see env:AUTH_PRIVATE_FILEPATH")
	}

	publicFilepath := os.Getenv("AUTH_PUBLIC_FILEPATH")
	if publicFilepath == "" {
		panic("auth middleware is not configured: see env:AUTH_PUBLIC_FILEPATH")
	}

	private, err := jwtAuth.GetECPrivateKeyFromFile(privateFilepath)
	if err != nil {
		panic(err)
	}

	public, err := jwtAuth.GetECPublicKeyFromFile(publicFilepath)
	if err != nil {
		panic(err)
	}

	authCli := jwtAuth.NewJwtAuth[User](jwtAuth.ES256, time.Hour*24, private, public, jwtAuth.JwtAuthConfig{
		Audience: Audience,
		Issuer:   Issuer,
		Subject:  Subject,
	})
	return middleware(authCli, DefaultVerificationFn)
}

func middleware(authCli *jwtAuth.JwtAuth[User], verificationFn VerificationFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			auth := tokenFromHeader(request)

			ok, err := verificationFn(auth, request.Context())
			if err != nil {
				return httperr.Respond(c, e.AuthortionErr, nil)
			}

			if !ok {
				user, _, err := authCli.Parse(auth)
				if err != nil {
					return httperr.Respond(c, e.LoginExpired, nil)
				}

				c.Set(userContextKey, user)
			}
			return next(c)
		}
	}
}

var DefaultVerificationFn = func(auth string, ctx context.Context) (bool, error) {
	_, ok := ctx.Value(BearerScopes).([]string)
	return !ok, nil
}
