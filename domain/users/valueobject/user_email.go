package valueobject

import (
	"blog/infra/e"
	"database/sql/driver"
	"strings"
)

type Email string

// NewEmail 新建email
func NewEmail(email string) (Email, error) {
	if len(strings.Split(email, "@")) != 2 {
		return "", nil
	}
	return Email(email), nil
}

// Scan Scanner接口 用于将数据库数据写入结构体
func (em *Email) Scan(src any) error {
	email, ok := src.([]byte)
	if !ok {
		return e.NewError(e.ScanSetErr, nil)
	}
	*em = Email(email)
	return nil
}

// Value Valuer接口 用于结构体数据写入数据库
func (em *Email) Value() (driver.Value, error) {
	return em, nil
}

// ToString 获取email
func (em *Email) ToString() string {
	return string(*em)
}

// GetNickName 获取email前半部分
func (em *Email) GetNickName() string {
	return strings.Split(string(*em), "@")[0]
}

// GetEmailDomainName 获取email域名
func (em *Email) GetEmailDomainName() string {
	return strings.Split(string(*em), "@")[1]
}
