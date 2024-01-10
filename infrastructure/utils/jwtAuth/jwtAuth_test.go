package jwtAuth_test

import (
	jwtAuth2 "blog/infrastructure/utils/jwtAuth"
	"github.com/google/uuid"
	"testing"
	"time"
)

type TestClaims struct {
	Name string
	Age  int
}

func TestJwtAuth(t *testing.T) {
	var (
		signKey   = []byte("test")
		verifyKey = []byte("test")
		exp       = time.Second
	)

	authorize, err := jwtAuth2.NewJwtAuthWithOption[TestClaims](
		"HS256",
		jwtAuth2.JwtAuthConfig{
			Issuer:    "Server",
			Subject:   "Test",
			Audience:  "User",
			NotBefore: 0,
		},
		signKey,
		verifyKey,
		jwtAuth2.JwtAuthModifyClaimsExpires(exp),
	)

	authorize.AddOption(jwtAuth2.JwtAuthModifyClaimsJwtID(
		func() string {
			return uuid.New().String()
		}),
		jwtAuth2.JwtAuthModifyClaimsIssuedAt,
	)

	if err != nil {
		t.Error(err)
		return
	}
	tokenStr, err := authorize.Encode(TestClaims{
		Name: "box",
		Age:  18,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tokenStr)

	claims, err := authorize.Decode(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims)

	claims2, meta, err := authorize.DecodeMeta(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims2)
	t.Logf("%v", meta)
}

func TestJwtAuthEC(t *testing.T) {
	var (
		err error
		exp = time.Second
	)

	publishKey, err := jwtAuth2.GetECPublicKeyFromFile("shared/key/es/public_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	privateKey, err := jwtAuth2.GetECPrivateKeyFromFile("shared/key/es/private_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	authorize, err := jwtAuth2.NewJwtAuthWithOption[TestClaims](
		"ES256",
		jwtAuth2.JwtAuthConfig{
			Issuer:   "Server",
			Subject:  "UserAuthorize",
			Audience: "User",
		},
		privateKey,
		publishKey,
		jwtAuth2.JwtAuthModifyClaimsExpires(exp),
	)

	if err != nil {
		t.Error(err)
		return
	}
	tokenStr, err := authorize.Encode(TestClaims{
		Name: "box",
		Age:  18,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tokenStr)

	claims, err := authorize.Decode(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims)

	claims2, meta, err := authorize.DecodeMeta(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims2)
	t.Logf("%v", meta)
}
