package jwtAuth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var signMap = map[string]jwt.SigningMethod{
	"HS256": jwt.SigningMethodHS256,
	"HS384": jwt.SigningMethodHS384,
	"HS512": jwt.SigningMethodHS512,
	"RS256": jwt.SigningMethodRS256,
	"RS384": jwt.SigningMethodRS384,
	"RS512": jwt.SigningMethodRS512,
	"ES256": jwt.SigningMethodES256,
	"ES384": jwt.SigningMethodES384,
	"ES512": jwt.SigningMethodES512,
	"PS256": jwt.SigningMethodPS256,
	"PS384": jwt.SigningMethodPS384,
	"PS512": jwt.SigningMethodPS512,
	"None":  jwt.SigningMethodNone,
}

var (
	SignMethodErr = errors.New("未实现的加密算法")
)

type JwtAuthOptionFunc func(claims *jwt.StandardClaims)

// JwtAuthConfig //
// 设置某一项即代表使用某一项，未设置的项不会出现在token中
type JwtAuthConfig struct {
	Issuer    string // 令牌的发行者，表示谁颁发了令牌
	Subject   string // 令牌的主题，通常是用户的唯一标识。表示令牌所代表的用户或实体
	Audience  string // 令牌的预期接收者，表示令牌的预期使用方
	NotBefore int64  // 生效时间，表示令牌在此时间之前不能被接受
}

type JwtAuth[Claims any] struct {
	alg        jwt.SigningMethod
	option     []JwtAuthOptionFunc
	signKey    any
	verifyKey  any
	metaClaims jwt.StandardClaims
}

func NewJwtAuth[Claims any](alg string, config JwtAuthConfig, signKey, verifyKey any) (*JwtAuth[Claims], error) {
	method, ok := signMap[alg]
	if !ok {
		return nil, SignMethodErr
	}
	return &JwtAuth[Claims]{
		alg:       method,
		option:    make([]JwtAuthOptionFunc, 0),
		signKey:   signKey,
		verifyKey: verifyKey,
		metaClaims: jwt.StandardClaims{
			Audience:  config.Audience,
			Issuer:    config.Issuer,
			NotBefore: config.NotBefore,
			Subject:   config.Subject,
		},
	}, nil
}

func NewJwtAuthWithOption[Claims any](alg string, config JwtAuthConfig, signKey, verifyKey any, fns ...JwtAuthOptionFunc) (*JwtAuth[Claims], error) {
	j, err := NewJwtAuth[Claims](alg, config, signKey, verifyKey)
	if err != nil {
		return nil, err
	}
	j.AddOption(fns...)
	return j, nil
}

func (j *JwtAuth[Claims]) AddOption(fns ...JwtAuthOptionFunc) {
	for _, fn := range fns {
		j.option = append(j.option, fn)
	}
}

func (j *JwtAuth[Claims]) Encode(claims Claims) (string, error) {
	return jwt.NewWithClaims(j.alg, j.getClaims(claims)).SignedString(j.signKey)
}

func (j *JwtAuth[Claims]) Decode(tokenStr string) (*Claims, error) {
	claims, _, err := j.DecodeMeta(tokenStr)
	return claims, err
}

func (j *JwtAuth[Claims]) DecodeMeta(tokenStr string) (*Claims, *jwt.StandardClaims, error) {
	var wrapedClaims = &wrapClaims[Claims]{}
	_, err := jwt.ParseWithClaims(tokenStr, wrapedClaims, j.getkeyFunc)
	if err != nil {
		return nil, nil, err
	}
	return &wrapedClaims.Payload, &wrapedClaims.StandardClaims, nil
}
func (j *JwtAuth[Claims]) getClaims(claims Claims) jwt.Claims {
	metaClaims := j.metaClaims
	for _, optionFn := range j.option {
		optionFn(&metaClaims)
	}
	return wrapClaims[Claims]{
		StandardClaims: metaClaims,
		Payload:        claims,
	}
}

func (j *JwtAuth[Claims]) getkeyFunc(token *jwt.Token) (interface{}, error) {
	return j.verifyKey, nil
}

type wrapClaims[Claims any] struct {
	jwt.StandardClaims
	Payload Claims
}
