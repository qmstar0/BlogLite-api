package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/qmstar0/BlogLite-api/config"
	"time"
)

func NewUserCliaims(id uint32, usertype, name string, duration time.Duration) UserClaims {
	now := time.Now()
	return UserClaims{
		ID:   id,
		Type: usertype,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			NotBefore: jwt.NewNumericDate(now),
			ID:        uuid.New().String(),
		},
	}
}

func Sign(claims UserClaims) (string, error) {
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.Cfg.AuthSecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Verify(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.AuthSecretKey), nil
	})

	if err != nil {
		return nil, errors.New("无法解析Token")
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效Token")
}
