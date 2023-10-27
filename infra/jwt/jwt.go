package jwt

import (
	"blog/infra/config"
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type Claims = jwt.StandardClaims

func init() {
	// 读取 RSA 私钥和公钥文件
	// PrivateKeyPath 和 PublicKeyPath 分别是你的 RSA 私钥和公钥的文件路径
	privateKeyBytes, err := os.ReadFile(config.Conf.Jwt.PrivateKeyPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read private key: %v", err))
	}

	publicKeyBytes, err := os.ReadFile(config.Conf.Jwt.PublicKeyPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read public key: %v", err))
	}

	// 解析 RSA 私钥和公钥
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse private key: %v", err))
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse public key: %v", err))
	}
}

// Sign 使用 RSA 私钥签名 JWT 令牌
func Sign(c jwt.Claims) (string, error) {
	// 使用HS256签名方法创建令牌
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, c)

	// 签名令牌并返回
	token, err := tokenClaims.SignedString(privateKey) // 使用jwtSecret进行签名
	return token, err
}

// ParseToken 使用 RSA 公钥验证和解析 JWT 令牌
func ParseToken(tokenString string, t jwt.Claims) (*jwt.Token, error) {
	// 解析JWT令牌，使用与签发时相同的签名方法
	token, err := jwt.ParseWithClaims(tokenString, t, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil // 使用与签发时相同的密钥
	})
	return token, err
}
