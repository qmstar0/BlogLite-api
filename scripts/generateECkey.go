package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateECDSAKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	// 选择椭圆曲线，这里使用P-256曲线
	curve := elliptic.P256()

	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// 获取公钥
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func savePrivateKeyToFile(privateKey *ecdsa.PrivateKey, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 使用PEM编码保存私钥
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	return pem.Encode(file, block)
}

func savePublicKeyToFile(publicKey *ecdsa.PublicKey, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 使用PEM编码保存公钥
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return pem.Encode(file, block)
}

func main() {
	// 生成ECDSA密钥对
	privateKey, publicKey, err := generateECDSAKeyPair()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 保存私钥到文件
	err = savePrivateKeyToFile(privateKey, "private_key.pem")
	if err != nil {
		fmt.Println("Error saving private key:", err)
		return
	}

	// 保存公钥到文件
	err = savePublicKeyToFile(publicKey, "public_key.pem")
	if err != nil {
		fmt.Println("Error saving public key:", err)
		return
	}

	fmt.Println("ECDSA key pair generated and saved successfully.")
}
