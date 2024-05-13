// Rxxxx Repository 仓储 数据、数据库、资源相关
// Pxxxx Params 外部参数相关
// Dxxxx Domain 领域相关
// Uxxxx Util 工具相关
// Axxxx Auth 认证相关

package e

import "fmt"

type StateCode struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Err  error  `json:"error,omitempty"`
}

func (e StateCode) Error() string {
	return e.Msg
}

func (e StateCode) WithMessage(msg string) StateCode {
	e.Msg = fmt.Sprintf("%s | %s", e.Msg, msg)
	return e
}

func (e StateCode) WithError(err error) StateCode {
	e.Err = err
	return e
}

var (
	NotImplement     = StateCode{Code: "?9999", Msg: "未知错误"}
	Successed        = StateCode{Code: "OK", Msg: ""}
	PErrInvalidParam = StateCode{Code: "P0101", Msg: "无效参数"}

	AErrUnauthortion       = StateCode{Code: "A0201", Msg: "未认证"}
	RErrResourceNotExists  = StateCode{Code: "R0101", Msg: "资源不存在"}
	RErrResourceExists     = StateCode{Code: "R0102", Msg: "资源已存在"}
	RErrDatabase           = StateCode{Code: "R0901", Msg: "未知错误"}
	DErrInvalidOperation   = StateCode{Code: "D0101", Msg: "无效操作"}
	UErrUtilMarkdownToHTML = StateCode{Code: "U0101", Msg: "无法处理Markdown文件"}
)
