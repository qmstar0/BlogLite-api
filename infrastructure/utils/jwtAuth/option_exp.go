package jwtAuth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JwtAuthModifyClaimsExpires(t time.Duration) JwtAuthOptionFunc {
	return func(claims *jwt.StandardClaims) {
		claims.ExpiresAt = time.Now().Add(t).Unix()
	}
}
