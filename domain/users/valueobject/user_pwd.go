package valueobject

import (
	"blog/infra/e"
	"database/sql/driver"
	"golang.org/x/crypto/bcrypt"
)

var (
	userPasswordCost = 12
)

type Passowrd string

// NewPassword 直接将传入值存入对象
func NewPassword(pwd string) (Passowrd, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), userPasswordCost)
	if err != nil {
		return "", err
	}
	return Passowrd(hashPwd), nil
}

// NewSafePassword 初始化时自动加密

// Scan Scanner接口 用于将数据库数据写入结构体
func (p *Passowrd) Scan(src any) error {
	pwd, ok := src.([]byte)
	if !ok {
		return e.NewError(e.ScanSetErr, nil)
	}
	*p = Passowrd(pwd)
	return nil
}

// Value Valuer接口 用于结构体数据写入数据库
func (p *Passowrd) Value() (driver.Value, error) {
	return p.ToString(), nil
}

// Check 验证密码
func (p *Passowrd) Check(pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p.ToString()), []byte(pwd)); err != nil {
		return false
	}
	return true
}

// ToString 获取hash password字符串
func (p *Passowrd) ToString() string {
	return string(*p)
}
