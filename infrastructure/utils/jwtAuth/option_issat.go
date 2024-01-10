package jwtAuth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JwtAuthModifyClaimsIssuedAt(claims *jwt.StandardClaims) {
	claims.IssuedAt = time.Now().Unix()
}
