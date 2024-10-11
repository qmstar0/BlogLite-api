package e

type ErrCode struct {
	Code    string
	Message string
}

func (e ErrCode) Error() string {
	return e.Message
}

func (e ErrCode) WithMessage(msg string) error {
	return ErrCode{Code: e.Code, Message: msg}
}

var (
	Unimplemented        = ErrCode{"UNIMPLEMENTED", ""}
	Unauthorised         = ErrCode{"UNAUTHORISED", ""}
	InternalServiceErr   = ErrCode{"INTERNAL_SERVICE_ERROR", ""}
	InvalidParameters    = ErrCode{"INVALID_PARAMETERS", ""}
	InvalidAction        = ErrCode{"INVALID_ACTION", ""}
	ResourceDoesNotExist = ErrCode{"RESOURCE_NOT_EXIST", "资源不存在或已删除"}
	PWDError             = ErrCode{"PWD_ERROR", "密码错误，请重新输入"}
)

func InvalidParametersError(info string) error {
	return InvalidParameters.WithMessage(info)
}

func InternalServiceError(info string) error {
	return InternalServiceErr.WithMessage(info)
}

func InvalidActionError(info string) error {
	return InvalidAction.WithMessage(info)
}

func UnauthorisedError(info string) error {
	return Unauthorised.WithMessage(info)
}
