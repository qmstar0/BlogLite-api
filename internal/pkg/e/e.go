// Rxxxx Repository 仓储 数据、数据库、资源相关
// Pxxxx Params 外部参数相关
// Dxxxx Domain 领域相关
// Uxxxx Util 工具相关
// Axxxx Auth 认证相关

package e

type StateCode struct {
	Code  string `json:"code,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Debug string `json:"debug,omitempty"`
}

func (e StateCode) Error() string {
	return e.Msg
}

func (e StateCode) WithMessage(msg string) StateCode {
	e.Msg = msg
	return e
}

func (e StateCode) WithError(err error) StateCode {
	e.Debug = err.Error()
	return e
}

var (
	NotImplement     = StateCode{Code: "E9999", Msg: "未知错误"}
	Successed        = StateCode{Code: "OK", Msg: ""}
	PErrInvalidParam = StateCode{Code: "P0101", Msg: "无效参数"}

	AErrUnauthortion       = StateCode{Code: "A0101", Msg: "未认证"}
	AErrWrongAuthortion    = StateCode{Code: "A0201", Msg: "认证错误"}
	AErrSignError          = StateCode{Code: "A0301", Msg: "签发Token出现错误"}
	RErrResourceNotExists  = StateCode{Code: "R0101", Msg: "资源不存在"}
	RErrResourceExists     = StateCode{Code: "R0102", Msg: "资源已存在"}
	RErrDatabase           = StateCode{Code: "R5901", Msg: "内部错误"}
	DErrInvalidOperation   = StateCode{Code: "D0101", Msg: "无效操作"}
	UErrUtilMarkdownToHTML = StateCode{Code: "U0101", Msg: "无法处理Markdown文件"}
)
