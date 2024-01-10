package authorize

import (
	jwtAuth2 "blog/infrastructure/utils/jwtAuth"
	"errors"
	"net/http"
	"strings"
	"time"
)

type AuthorizeClaims struct {
	Uid int
}

var (
	Authorize *jwtAuth2.JwtAuth[AuthorizeClaims]
	timeExp   = time.Hour
)

func init() {
	var err error

	publishKey, err := jwtAuth2.GetECPublicKeyFromFile("shared/key/es/public_key.pem")
	if err != nil {
		panic(err)
	}
	privateKey, err := jwtAuth2.GetECPrivateKeyFromFile("shared/key/es/private_key.pem")
	if err != nil {
		panic(err)
	}
	Authorize, err = jwtAuth2.NewJwtAuthWithOption[AuthorizeClaims](
		"ES256",
		jwtAuth2.JwtAuthConfig{
			Issuer:   "Server",
			Subject:  "UserAuthorize",
			Audience: "User",
		},
		privateKey,
		publishKey,
		jwtAuth2.JwtAuthModifyClaimsExpires(timeExp),
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
	NoAuthorizeInformationErr   = errors.New("没有Authorize相关信息")
)
