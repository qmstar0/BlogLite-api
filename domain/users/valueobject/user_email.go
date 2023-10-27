package valueobject

import (
	er "blog/infra/e"
	"database/sql/driver"
	"strings"
)

type Email struct {
	Data            string `json:"Data"`
	NickName        string `json:"nick_name"`
	EmailDomainName string `json:"email_domain_name"`
}

// NewEmail 新建email
func NewEmail(email string) (Email, error) {
	s := strings.Split(email, "@")
	if len(s) != 2 {
		return Email{}, er.NewError(er.EmailFormatErr, nil)
	}

	return Email{
		Data:            email,
		NickName:        s[0],
		EmailDomainName: s[1],
	}, nil
}

// Scan Scanner接口 用于将数据库数据写入结构体
func (e Email) Scan(src any) error {
	email := string(src.([]byte))
	s := strings.Split(email, "@")
	if len(s) != 2 {
		return er.NewError(er.EmailFormatErr, nil)
	}
	e.Data = email
	e.NickName = s[0]
	e.EmailDomainName = s[1]
	return nil
}

// Value Valuer接口 用于结构体数据写入数据库
func (e Email) Value() (driver.Value, error) {
	return e.ToString(), nil
}

// ToString 获取email
func (e Email) ToString() string {
	return e.Data
}

// GetNickName 获取email前半部分
func (e Email) GetNickName() string {
	return e.NickName
}

// GetEmailDomainName 获取email域名
func (e Email) GetEmailDomainName() string {
	return e.EmailDomainName
}
