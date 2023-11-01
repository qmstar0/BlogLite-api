package token

import (
	"blog/infra/e"
	"blog/infra/jwt"
	"errors"
	"github.com/google/uuid"
	"time"
)

const (
	day          = time.Hour * 24
	ExpAuthToken = 180 * day
	AuthAudience = "Auth"
)

type AuthClaims struct {
	Email string
	Uid   string
	Name  string
	Role  uint
	jwt.Claims
}

func NewAuthToken(uid, email, name string, role uint) (string, error) {
	t := time.Now()

	claims := AuthClaims{
		Email: email,
		Uid:   uid,
		Name:  name,
		Role:  role,
		Claims: jwt.Claims{
			Audience:  AuthAudience,
			ExpiresAt: t.Add(ExpAuthToken).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  t.Unix(),
			Issuer:    "Auth",
			Subject:   "AuthToken",
		},
	}
	sign, err := jwt.Sign(claims)
	if err != nil {
		return "", e.NewError(e.JwtSignErr, err)
	}
	return sign, nil
}

func ParseAuthToken(t string) (*AuthClaims, error) {
	token, err := jwt.ParseToken(t, &AuthClaims{})
	if err != nil {
		return nil, e.NewError(e.JwtParseErr, err)
	}
	claims, ok := token.Claims.(*AuthClaims)
	if !ok {
		return nil, e.NewError(e.JwtParseErr, errors.New("解析为Claims结构体时错误"))
	}
	return claims, nil
}
