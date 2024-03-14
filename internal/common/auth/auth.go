package auth

import (
	"blog/pkg/jwtAuth"
	"crypto/ecdsa"
	"fmt"
	"os"
	"sync"
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

var (
	private *ecdsa.PrivateKey
	public  *ecdsa.PublicKey
)

func init() {
	var err error
	privateFilepath := os.Getenv("AUTH_PRIVATE_FILEPATH")
	if privateFilepath == "" {
		panic("auth middleware is not configured: see env:AUTH_PRIVATE_FILEPATH")
	}

	publicFilepath := os.Getenv("AUTH_PUBLIC_FILEPATH")
	if publicFilepath == "" {
		panic("auth middleware is not configured: see env:AUTH_PUBLIC_FILEPATH")
	}

	private, err = jwtAuth.GetECPrivateKeyFromFile(privateFilepath)
	if err != nil {
		panic(err)
	}

	public, err = jwtAuth.GetECPublicKeyFromFile(publicFilepath)
	if err != nil {
		panic(err)
	}
}

func NewAuthClient[E any]() *jwtAuth.JwtAuth[E] {
	name := fmt.Sprintf("%T", new(E))
	Cli, ok := authClientMap.LoadOrStore(name, &jwtAuth.JwtAuth[E]{})
	if ok {
		return Cli.(*jwtAuth.JwtAuth[E])
	}
	authCli := jwtAuth.NewJwtAuth[E](jwtAuth.ES256, time.Hour*24, private, public, jwtAuth.JwtAuthConfig{
		Audience: Audience,
		Issuer:   Issuer,
		Subject:  Subject,
	})
	authClientMap.Store(name, authCli)
	return authCli
}

var authClientMap sync.Map
