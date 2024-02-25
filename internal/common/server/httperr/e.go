package httperr

type StateCode int

// Successed 成功
const Successed StateCode = 0

// NotImplementedErr 临时占位错误
const NotImplementedErr StateCode = 99999

// 输入输出错误
const (
	InputErr StateCode = iota + 21010
	InvalidParam
	MissingParam
	ResourceDoesNotExist
	EmailFormatErr
)

// 安全和鉴权错误
const (
	SuspiciousRequest StateCode = iota + 31010
	AuthortionErr
	PermissionDenied
	LoginExpired
	LoginRequired
)

// ServeErr 后台错误
const ServeErr StateCode = iota + 40010

// 通信错误
const (
	SendErr StateCode = iota + 51010
	EmailSendErr
)

// 领域服务或事件错误
const (
	DomainErr StateCode = iota + 60010
	DomainEventDataTypeErr
	CommandHandlerErr
)

var errMsg = map[StateCode]string{
	Successed: "Successed",

	SendErr:              "Publish Err",
	ServeErr:             "Serve Err",
	AuthortionErr:        "Authortion Err",
	InputErr:             "Input Err",
	ResourceDoesNotExist: "Resource Does Not Exist",
	DomainErr:            "Domain Err",

	InvalidParam:      "Invalid Param",
	MissingParam:      "Missing Param",
	EmailFormatErr:    "email format httpError",
	PermissionDenied:  "Permission Denied",
	LoginRequired:     "Login Required",
	LoginExpired:      "Login Expired",
	SuspiciousRequest: "Suspicious Request",

	EmailSendErr: "Email Publish Err",

	DomainEventDataTypeErr: "Domain Event Data Type Err",
	CommandHandlerErr:      "Command Handler Err",
}

func (s StateCode) Error() string {
	msg, _ := errMsg[s]
	return msg
}
