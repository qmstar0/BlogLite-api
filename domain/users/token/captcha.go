package token

import (
	"blog/infra/config"
	"blog/infra/e"
	"blog/infra/jwt"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	ExpCaptchaToken = time.Minute * 3
	CaptchaAudience = "Captcha"
)

type CaptchaClaims struct {
	Email      string
	HashCaptch string
	jwt.Claims
}

func (c *CaptchaClaims) Verify(email, captcha string) error {
	hash, err := getHash(captcha)
	if err != nil {
		return err
	}
	if c.Email != email || c.HashCaptch != hash {
		return e.NewError(e.TokenVerifyErr, nil)
	}
	return nil
}

func NewCaptchaToken(email, Captcha string) (string, error) {
	t := time.Now()
	hashCaptcha, _ := getHash(Captcha)
	claims := CaptchaClaims{
		Email:      email,
		HashCaptch: hashCaptcha,
		Claims: jwt.Claims{
			Audience:  CaptchaAudience,
			ExpiresAt: t.Add(ExpCaptchaToken).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  t.Unix(),
			Issuer:    "Captcha",
			Subject:   "CaptchaToken",
		},
	}
	sign, err := jwt.Sign(claims)
	if err != nil {
		return "", e.NewError(e.JwtSignErr, err)
	}
	return sign, nil
}

func ParseCaptchaToken(t string) (*CaptchaClaims, error) {
	token, err := jwt.ParseToken(t, &CaptchaClaims{})
	if err != nil {
		fmt.Println(err)
		return nil, e.NewError(e.JwtParseErr, err)
	}
	claims, ok := token.Claims.(*CaptchaClaims)
	if !ok {
		return nil, e.NewError(e.JwtParseErr, errors.New("解析为Claims结构体时错误"))
	}
	return claims, nil
}

func getHash(c string) (string, error) {
	b := []byte(c + config.Conf.User.HashCaptchaSalt)
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		return "", err
	}
	hashBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString, err
}
