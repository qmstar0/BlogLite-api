package jwt

import (
	"blog/infra/config"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type MapClaims = jwt.MapClaims

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

// iss（Issuer）:含义：令牌的发行者，表示谁颁发了令牌。
// 示例："iss": "your-issuer"
// sub（Subject）:含义：令牌的主题，通常是用户的唯一标识。表示令牌所代表的用户或实体。
// 示例："sub": "users-id"
// aud（Audience）:含义：令牌的预期接收者，表示令牌的预期使用方。
// 示例："aud": "your-audience"
// exp（Expiration Time）:含义：令牌的过期时间，表示令牌的有效期截止时间。
// 示例："exp": 1609459200（Unix 时间戳）
// nbf（Not Before）:含义：生效时间，表示令牌在此时间之前不能被接受。
// 示例："nbf": 1609459200（Unix 时间戳）
// iat（Issued At）:含义：令牌的发行时间，表示令牌被创建的时间。
// 示例："iat": 1609459200（Unix 时间戳）
// jti（JWT ID）:含义：JWT 的唯一标识符，用于防止 JWT 重放攻击。
// 示例："jti": "unique-id"

// Sign 使用 RSA 私钥签名 JWT 令牌
func Sign(data map[string]any, expTime time.Duration) (string, error) {
	claims := jwt.MapClaims(data)
	t := time.Now()
	claims["exp"] = t.Add(expTime).Unix()
	claims["iat"] = t.Unix()
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// 签名令牌并返回
	token, err := tokenClaims.SignedString(privateKey) // 使用jwtSecret进行签名
	return token, err
}

// ParseToken 使用 RSA 公钥验证和解析 JWT 令牌
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	// 解析JWT令牌，使用与签发时相同的签名方法
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil // 使用与签发时相同的密钥
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	if !parsedToken.Valid {
		return jwt.MapClaims{}, errors.New("token is invalid")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errors.New("an error occurred in claims, ok := parsedToken.Claims.(jwt.MapClaims)")
	}
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return jwt.MapClaims{}, errors.New("token is Expired")
	}
	return claims, nil
}
