// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ports

const (
	BearerScopes = "bearer.Scopes"
)

// UserLoginJSONBody defines parameters for UserLogin.
type UserLoginJSONBody struct {
	// Email 有效邮箱
	Email string `json:"email"`

	// Password 加密密码
	Password string `json:"password"`
}

// UserLogoutJSONBody defines parameters for UserLogout.
type UserLogoutJSONBody = map[string]interface{}

// UserRegisterJSONBody defines parameters for UserRegister.
type UserRegisterJSONBody struct {
	// Email 有效邮箱
	Email string `json:"email"`

	// Password 加密密码
	Password string `json:"password"`

	// Token token
	Token string `json:"token"`
}

// UserApplyRegisterJSONBody defines parameters for UserApplyRegister.
type UserApplyRegisterJSONBody = map[string]interface{}

// UserApplyRegisterParams defines parameters for UserApplyRegister.
type UserApplyRegisterParams struct {
	Email *string `form:"email,omitempty" json:"email,omitempty"`
}

// ResetPasswordJSONBody defines parameters for ResetPassword.
type ResetPasswordJSONBody struct {
	// Password 新密码
	Password string `json:"password"`
}

// ModifyUserNameJSONBody defines parameters for ModifyUserName.
type ModifyUserNameJSONBody struct {
	// Name 新用户名
	Name string `json:"name"`
}

// UserVerifyJSONBody defines parameters for UserVerify.
type UserVerifyJSONBody struct {
	Password string `json:"password"`
}

// UserLoginJSONRequestBody defines body for UserLogin for application/json ContentType.
type UserLoginJSONRequestBody UserLoginJSONBody

// UserLogoutJSONRequestBody defines body for UserLogout for application/json ContentType.
type UserLogoutJSONRequestBody = UserLogoutJSONBody

// UserRegisterJSONRequestBody defines body for UserRegister for application/json ContentType.
type UserRegisterJSONRequestBody UserRegisterJSONBody

// UserApplyRegisterJSONRequestBody defines body for UserApplyRegister for application/json ContentType.
type UserApplyRegisterJSONRequestBody = UserApplyRegisterJSONBody

// ResetPasswordJSONRequestBody defines body for ResetPassword for application/json ContentType.
type ResetPasswordJSONRequestBody ResetPasswordJSONBody

// ModifyUserNameJSONRequestBody defines body for ModifyUserName for application/json ContentType.
type ModifyUserNameJSONRequestBody ModifyUserNameJSONBody

// UserVerifyJSONRequestBody defines body for UserVerify for application/json ContentType.
type UserVerifyJSONRequestBody UserVerifyJSONBody
