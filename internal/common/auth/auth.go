package auth

import (
	"blog/pkg/jwtAuth"
	httperr2 "common/server/httperr"
	"context"
	"errors"
	"net/http"
	"strings"
)

type User struct {
	Name  string
	Email string
	ID    string
	Role  int16
}

type AuthorizationHttpMiddleware struct {
	AuthClient *jwtAuth.JwtAuth[User]
}

func (a AuthorizationHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			httperr2.Respond(w, httperr2.LoginRequired, "please sign in")
			return
		}

		user, _, err := a.AuthClient.Parse(bearerToken)
		if err != nil {
			httperr2.Respond(w, httperr2.LoginExpired, err.Error())
			return
		}

		// it's always a good idea to use custom type as context value (in this case ctxKey)
		// because nobody from the outside of the package will be able to override/read this value
		ctx = context.WithValue(ctx, userContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a AuthorizationHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

type ctxKey int

const userContextKey ctxKey = iota

var NoUserInContextError = errors.New("no user in context")

func GetUserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
