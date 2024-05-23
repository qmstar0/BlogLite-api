package auth_test

import (
	auth2 "go-blog-ddd/pkg/auth"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	authenticator := auth2.NewJWTAuthenticator([]byte("test"), "ADMIN PERMISSIONS", "于野|探索日志", "ADMIN", "AUTHOR")
	signStr, claims, err := authenticator.Sign(time.Second * 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(claims)
	t.Log(signStr)

	_, err = authenticator.Parse(signStr)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 1)
	_, err = authenticator.Parse(signStr)
	if err == nil {
		t.Fatal("token超时后仍可用")
	}

}

func TestAPIKey(t *testing.T) {
	const ApiKeyLength = 32
	authenticator := auth2.NewAPIKeyAuthenticator(ApiKeyLength)
	token, err := authenticator.Sign()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	err = authenticator.Check(token)
	if err != nil {
		t.Fatal(err)
	}
}
