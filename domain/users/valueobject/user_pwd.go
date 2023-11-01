package valueobject

import (
	"database/sql/driver"
	"golang.org/x/crypto/bcrypt"
)

var (
	userPasswordCost = 12
)

type Passowrd struct {
	HashData string `json:"hash_data"`
}

// NewPassword 直接将传入值存入对象
func NewPassword(pwd string) Passowrd {
	return Passowrd{HashData: pwd}
}

// NewSafePassword 初始化时自动加密
func NewSafePassword(pwd string) (Passowrd, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), userPasswordCost)
	if err != nil {

		return Passowrd{}, err
	}
	return Passowrd{HashData: string(hashPwd)}, nil
}

// Scan Scanner接口 用于将数据库数据写入结构体
func (p Passowrd) Scan(src any) error {
	p.HashData = string(src.([]byte))
	return nil
}

// Value Valuer接口 用于结构体数据写入数据库
func (p Passowrd) Value() (driver.Value, error) {
	return p.ToString(), nil
}

// VerifyPassword 验证密码
func (p Passowrd) VerifyPassword(pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p.ToString()), []byte(pwd)); err != nil {
		return false
	}
	return true
}

// ToString 获取hash password字符串
func (p Passowrd) ToString() string {
	return p.HashData
}
