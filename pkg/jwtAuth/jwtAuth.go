package jwtAuth

import (
	"bytes"
	"encoding/gob"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JwtAuthConfig //
// 设置某一项即代表使用某一项，未设置的项不会出现在token中
//type JwtAuthConfig struct {
//	Issuer    string // 令牌的发行者，表示谁颁发了令牌
//	Subject   string // 令牌的主题，通常是用户的唯一标识。表示令牌所代表的用户或实体
//	Audience  string // 令牌的预期接收者，表示令牌的预期使用方
//	NotBefore int64  // 生效时间，表示令牌在此时间之前不能被接受
//}

var (
	ES256 = jwt.SigningMethodES256
	RS256 = jwt.SigningMethodRS256
	HS256 = jwt.SigningMethodHS256
)

type JwtAuthConfig struct {
	Audience string
	Issuer   string
	Subject  string
}

type JwtAuth[Claims any] struct {
	alg        jwt.SigningMethod
	exp        time.Duration
	iss        string
	signKey    any
	verifyKey  any
	metaClaims jwt.StandardClaims
}

func NewJwtAuth[Claims any](alg jwt.SigningMethod, exp time.Duration, signKey, verifyKey any, config JwtAuthConfig) *JwtAuth[Claims] {
	return &JwtAuth[Claims]{
		alg:       alg,
		exp:       exp,
		signKey:   signKey,
		verifyKey: verifyKey,
		metaClaims: jwt.StandardClaims{
			Issuer:   config.Issuer,
			Subject:  config.Subject,
			Audience: config.Audience,
		},
	}
}

func (j *JwtAuth[Claims]) Sign(claims Claims) (string, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(claims)
	if err != nil {
		return "", err
	}
	return jwt.NewWithClaims(j.alg, j.wrapClaims(buf.Bytes())).SignedString(j.signKey)
}

func (j *JwtAuth[Claims]) Parse(tokenStr string) (Claims, *jwt.StandardClaims, error) {
	var wrapedClaims = new(wrapClaims)
	var claims Claims

	_, err := jwt.ParseWithClaims(tokenStr, wrapedClaims, func(token *jwt.Token) (interface{}, error) {
		return j.verifyKey, nil
	})

	if err != nil {
		return claims, nil, err
	}
	err = gob.NewDecoder(bytes.NewBuffer(wrapedClaims.Payload)).Decode(&claims)
	if err != nil {
		return claims, nil, err
	}
	return claims, &wrapedClaims.StandardClaims, nil
}

func (j *JwtAuth[Claims]) wrapClaims(payload []byte) jwt.Claims {
	now := time.Now()
	nowUnix := now.Unix()
	metaClaims := j.metaClaims
	metaClaims.ExpiresAt = now.Add(j.exp).Unix()
	metaClaims.NotBefore = nowUnix
	metaClaims.IssuedAt = nowUnix
	return wrapClaims{
		StandardClaims: metaClaims,
		Payload:        payload,
	}
}

type wrapClaims struct {
	jwt.StandardClaims
	Payload []byte
}
