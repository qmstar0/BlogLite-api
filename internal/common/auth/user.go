package auth

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	ID   uint32
	Type string
	Name string
	jwt.RegisteredClaims
}

func (u UserClaims) IsAnyUserType(ss ...string) bool {
	for _, s := range ss {
		if s == u.Type {
			return true
		}
	}
	return false
}
