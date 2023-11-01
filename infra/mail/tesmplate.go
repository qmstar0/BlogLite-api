package mail

import (
	"bytes"
	"html/template"
)

const (
	TemplatePathForVerifyCode = "shared/templates/email_verifycode.html"
)

type templateBase struct {
	templateFilePath string
	buf              *bytes.Buffer
}

type ForVerifyCode struct {
	templateBase
	Email string
	Code  string
}

func NewTemplateForVerifyCode(email, code string) ForVerifyCode {
	t := ForVerifyCode{
		Email: email,
		Code:  code,
		templateBase: templateBase{
			templateFilePath: TemplatePathForVerifyCode,
			buf:              &bytes.Buffer{},
		},
	}
	temp, _ := template.ParseFiles(t.templateFilePath)
	_ = temp.Execute(t.buf, t)
	return t

}

func (t *templateBase) ToString() string {
	return t.buf.String()
}
