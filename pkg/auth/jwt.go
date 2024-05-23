package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var genID = func() string {
	return uuid.New().String()
}

type JWTAuthenticator struct {
	key    any
	claims jwt.RegisteredClaims
}

func NewJWTAuthenticator(key any, subject, issuer string, audience ...string) *JWTAuthenticator {
	return &JWTAuthenticator{
		key: key,
		claims: jwt.RegisteredClaims{
			Issuer:   issuer,
			Subject:  subject,
			Audience: audience,
		},
	}
}

func (a *JWTAuthenticator) Sign(duration time.Duration) (string, jwt.Claims, error) {
	now := time.Now()
	a.claims.IssuedAt = jwt.NewNumericDate(now)
	a.claims.NotBefore = jwt.NewNumericDate(now)
	a.claims.ExpiresAt = jwt.NewNumericDate(now.Add(duration))
	a.claims.ID = genID()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.claims)
	signedString, err := token.SignedString(a.key)
	if err != nil {
		return "", nil, err
	}
	return signedString, a.claims, nil
}

func (a *JWTAuthenticator) Parse(tokenStr string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, a.keyFn)
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func (a *JWTAuthenticator) keyFn(_ *jwt.Token) (interface{}, error) {
	return a.key, nil
}
