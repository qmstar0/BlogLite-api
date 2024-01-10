package jwtAuth

import (
	"github.com/dgrijalva/jwt-go"
)

func JwtAuthModifyClaimsJwtID(fn func() string) JwtAuthOptionFunc {
	return func(claims *jwt.StandardClaims) {
		claims.Id = fn()
	}
}
