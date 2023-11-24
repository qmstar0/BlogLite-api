package dto

//验证模型

// Captcha 验证请求
type Captcha struct {
	Email   string `json:"email" form:"email"  validate:"required,email"`
	Captcha string `json:"captcha" form:"captcha" validate:"required,len=6"`
	Token   string `json:"token" form:"token" validate:"required"`
}
