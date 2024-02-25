package jwtAuth_test

import (
	"blog/pkg/jwtAuth"
	"bytes"
	"encoding/gob"
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

	buffer := bytes.NewBuffer(make([]byte, 0, 4096))

	err := gob.NewEncoder(buffer).Encode(TestClaims{
		Name: "QMstar",
		Age:  18,
	})

	if err != nil {
		t.Fatal(err)
	}

	authorize := jwtAuth.NewJwtAuth[TestClaims](
		jwtAuth.HS256,
		exp,
		signKey,
		verifyKey,
		jwtAuth.JwtAuthConfig{},
	)

	tokenStr, err := authorize.Sign(TestClaims{
		Name: "QMstar",
		Age:  20,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tokenStr)

	claims, meta, err := authorize.Parse(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(claims, meta)
}

func TestJwtAuthEC(t *testing.T) {
	var (
		err error
		exp = time.Second
	)

	publishKey, err := jwtAuth.GetECPublicKeyFromFile("shared/key/es/public_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	privateKey, err := jwtAuth.GetECPrivateKeyFromFile("shared/key/es/private_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	authorize := jwtAuth.NewJwtAuth[TestClaims](
		jwtAuth.ES256,
		exp,
		privateKey,
		publishKey,
		jwtAuth.JwtAuthConfig{},
	)

	if err != nil {
		t.Error(err)
		return
	}
	tokenStr, err := authorize.Sign(TestClaims{
		Name: "box",
		Age:  18,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tokenStr)

	claims, meta, err := authorize.Parse(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims, meta)
}
