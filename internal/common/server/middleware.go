package server

import (
	"blog/pkg/jwtAuth"
	"categorys/ports"
	"common/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"os"
	"time"
)

func setupMiddleware(router chi.Router) {
	router.Use(middleware.Recoverer)

	//AddAuthMiddleware(router)
}

func AuthMiddleware() ports.MiddlewareFunc {
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
	authCli := jwtAuth.NewJwtAuth[auth.User](jwtAuth.ES256, time.Hour*24, private, public, jwtAuth.JwtAuthConfig{
		Audience: "User",
		Issuer:   "AuthorizationHttpMiddleware",
		Subject:  "UserToken",
	})

	return auth.AuthorizationHttpMiddleware{
		AuthClient: authCli,
	}.Middleware
}
