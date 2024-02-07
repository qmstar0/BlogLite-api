package authorize

import (
	"blog/infrastructure/utils/jwtAuth"
	"errors"
	"net/http"
	"strings"
	"time"
)

type AuthorizeClaims struct {
	Uid int
}

var (
	Authorize *jwtAuth.JwtAuth[AuthorizeClaims]
	timeExp   = time.Hour * 48
)

func init() {
	var err error

	publishKey, err := jwtAuth.GetECPublicKeyFromFile("shared/key/es/public_key.pem")
	if err != nil {
		panic(err)
	}
	privateKey, err := jwtAuth.GetECPrivateKeyFromFile("shared/key/es/private_key.pem")
	if err != nil {
		panic(err)
	}
	Authorize, err = jwtAuth.NewJwtAuthWithOption[AuthorizeClaims](
		"ES256",
		jwtAuth.JwtAuthConfig{
			Issuer:   "Server",
			Subject:  "UserAuthorize",
			Audience: "User",
		},
		privateKey,
		publishKey,
		jwtAuth.JwtAuthModifyClaimsExpires(timeExp),
	)
	if err != nil {
		panic(err)
	}
}

func ParseToClaims(r *http.Request) (*AuthorizeClaims, error) {
	token := getTokenFromAuthorization(r)
	if token == "" {
		return nil, NoAuthorizeInformationErr
	}
	claims, err := Authorize.Decode(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func SignFromClaims(claims AuthorizeClaims) (string, error) {
	return Authorize.Encode(claims)
}

func getTokenFromAuthorization(r *http.Request) string {
	authTokenStr := r.Header.Get("Authorization")
	splitStr := strings.Split(authTokenStr, " ")
	if len(splitStr) >= 2 {
		return splitStr[1]
	}
	return ""
}

var (
	PermissionVerificationError = errors.New("权限校验错误")
	NoAuthorizeInformationErr   = errors.New("authorize header missing")
)
