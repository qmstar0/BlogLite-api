package utils

import (
	"blog/infra/config"
	"crypto/sha256"
	"encoding/hex"
)

func GetSaltHash(c string) (string, error) {
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
