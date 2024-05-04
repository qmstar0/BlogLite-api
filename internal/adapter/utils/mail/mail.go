package mail

import (
	"go-blog-ddd/infra/config"
	"gopkg.in/gomail.v2"
	"os"
)

type Mail struct {
	name  string
	email string
	code  string
	host  string
	port  int
}

var (
	PostMan *Mail
)

func init() {
	if config.Conf.Mail.Email == "" {
		panic("Email is not configured: see config.yml")
	}
	mailCode := os.Getenv("BLOG_MAIL_PASSWORD")
	if mailCode == "" {
		panic("Email authorization related is not configured: see env:BLOG_MAIL_PASSWORD")
	}

	if config.Conf.Smtp.Host == "" || config.Conf.Smtp.Port == 0 {
		panic("smtp related not configured: see config.yml")
	}

	PostMan = &Mail{
		name:  config.Conf.Mail.Name,
		email: config.Conf.Mail.Email,
		code:  mailCode,
		host:  config.Conf.Smtp.Host,
		port:  config.Conf.Smtp.Port,
	}
}

// NewMail 新建邮件
func (m *Mail) NewMail() *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", m.email, m.name)
	return msg
}

// NewDialer 新建连接
func (m *Mail) NewDialer() *gomail.Dialer {
	dialer := gomail.NewDialer(m.host, m.port, m.email, m.code)
	return dialer
}
