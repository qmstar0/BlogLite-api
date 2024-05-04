package e

type StateCode struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e StateCode) Error() string {
	return e.Message
}

var NotImplemented = func(msg string) StateCode { return StateCode{Code: "?9999", Message: msg} }
var Successed = StateCode{"OK", ""}

var (
	InvalidParam         = StateCode{"P0101", "无效参数"}
	Unauthortion         = StateCode{"P0201", "未认证"}
	LoginRequired        = StateCode{"P0301", "必须登录"}
	LoginExpired         = StateCode{"P0302", "登录过期"}
	ResourceDoesNotExist = StateCode{"D0101", "资源不存在"}
	UserDoesNotExist     = StateCode{"D0201", "用户不存在"}

	ResourceAlreadyExists = StateCode{"D0102", "资源已存在"}
	UserAlreadyExists     = StateCode{"D0202", "用户已存在"}
)
