package jwtAuth

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func GetECPrivateKeyFromFile(p string) (*ecdsa.PrivateKey, error) {
	data, err := getFileData(p)
	if err != nil {
		return nil, err
	}
	return jwt.ParseECPrivateKeyFromPEM(data)
}

func GetECPublicKeyFromFile(p string) (*ecdsa.PublicKey, error) {
	data, err := getFileData(p)
	if err != nil {
		return nil, err
	}
	return jwt.ParseECPublicKeyFromPEM(data)
}

func GetRSAPrivateKeyFromFile(p string) (*rsa.PrivateKey, error) {
	data, err := getFileData(p)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(data)
}

func GetRSAPublicKeyFromFile(p string) (*rsa.PublicKey, error) {
	data, err := getFileData(p)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(data)
}

func getFileData(p string) ([]byte, error) {
	return os.ReadFile(p)
}
