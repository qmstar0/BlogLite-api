package utils

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	digits  = "0123456789"
)

var (
	random *rand.Rand
)

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomStr 生成随机字符
func RandomStr(length int) string {
	// 创建一个字符切片，用于存储生成的随机字符
	result := make([]byte, length)

	// 生成随机字符
	for i := 0; i < length; i++ {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

// RandomNum 生成随机数字
func RandomNum(length int) string {
	result := make([]byte, length)
	// 生成随机数字
	for i := 0; i < length; i++ {
		result[i] = digits[random.Intn(len(digits))]
	}

	return string(result)
}
