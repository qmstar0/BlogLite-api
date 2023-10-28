package response

import (
	"blog/infra/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	successCode = 0
)

// 响应模型
type response struct {
	Code    uint32 `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

// Api api响应
type Api struct {
	C *gin.Context
}

func (a Api) Success(data any) {
	a.C.JSON(http.StatusOK, response{
		Code:    successCode,
		Data:    data,
		Message: "Successed",
	})
	a.C.Abort()
}

// ValidateFailResp 400 参数验证失败
func (a Api) ValidateFailResp(err error) {
	a.failResponse(http.StatusBadRequest, err)
}

// TokenFailResp 401 Token验证失败
func (a Api) TokenFailResp(err error) {
	a.failResponse(http.StatusUnauthorized, err)
}

// UnauthorizedFailResp 401 没有权限
func (a Api) UnauthorizedFailResp(err error) {
	a.failResponse(http.StatusUnauthorized, err)
}

// ForbiddenFailResp 403 可疑请求
func (a Api) ForbiddenFailResp(err error) {
	a.failResponse(http.StatusForbidden, err)
}

// NotFoundFailResp 200 没有找到
func (a Api) NotFoundFailResp(err error) {
	a.failResponse(http.StatusOK, err)
}

// NotAcceptableFailResp 406 不允许的方法
func (a Api) NotAcceptableFailResp(err error) {
	a.failResponse(http.StatusNotAcceptable, err)
}

// ConflictFailResp 409 资源冲突
func (a Api) ConflictFailResp(err error) {
	a.failResponse(http.StatusConflict, err)
}

// UnavailableFailResp 503 业务逻辑失败
func (a Api) UnavailableFailResp(err error) {
	a.failResponse(http.StatusServiceUnavailable, err)
}

// failResponse 标准失败
func (a Api) failResponse(httpCode int, err error) {
	code, message := e.ParseErr(err)
	a.C.JSON(httpCode, response{
		Code:    code,
		Data:    nil,
		Message: message,
	})
	a.C.Abort()
}
