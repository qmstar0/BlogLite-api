package dto

//验证模型

// Captcha 验证请求
type Captcha struct {
	Email   string `form:"email"  validate:"required,email"`
	Captcha string `form:"captcha" validate:"required"`
	Token   string `form:"token" validate:"required"`
}
