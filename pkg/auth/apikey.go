package auth

import (
	"crypto/rand"
	"encoding/base64"
	"go-blog-ddd/internal/pkg/e"
)

type APIKeyAuthenticator struct {
	key    string
	length int
}

func NewAPIKeyAuthenticator(length int) *APIKeyAuthenticator {
	a := &APIKeyAuthenticator{
		key:    "",
		length: length,
	}
	_, err := a.Sign()
	if err != nil {
		panic(err)
	}
	return a
}

func (a *APIKeyAuthenticator) Sign() (token string, err error) {
	b := make([]byte, a.length)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	defer func() {
		a.key = token
	}()
	return base64.URLEncoding.EncodeToString(b), nil
}

func (a *APIKeyAuthenticator) Check(tokenStr string) error {
	if a.key != tokenStr {
		return e.AErrWrongAuthortion
	}
	return nil
}
