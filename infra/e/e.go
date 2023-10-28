package e

import (
	"blog/infra/logging"
	"errors"
	"fmt"
)

var logger *logging.Log

// Successed 成功
const Successed = 0

// NotImplementedErr 临时占位错误
const NotImplementedErr = 99999

// 输入输出错误
const (
	InputErr = iota + 21010
	InvalidParam
	MissingParam
	InvalidUpdate
	InvalidDelete
	ItemNotExist
	UserDuplicateCreationErr
	DuplicatePasswordErr
	EmailFormatErr
	PermissionDenied
)

// 安全和鉴权错误
const (
	SafeErr = iota + 31010
	JwtSignErr
	JwtParseErr
	TokenVerifyErr
	TokenExpiredErr
	SuspiciousRequest
)

// 后台错误
const (
	ServeErr = iota + 40010
	GenFilenameErr
	MarkdownTOHTMLErr
	PwdEncryptionErr
)

// 通信错误
const (
	SendErr = iota + 51010
	EmailSendErr
)

// 领域服务或事件错误
const (
	DomainErr = iota + 60010
	DomainEventDataTypeErr
)

// 数据持久化错误
const (
	DataPersistenceErr = iota + 80010
	DBCreateErr
	DBUpdateErr
	DBDeleteErr
	DBFindErr
	CacheCreateErr
	CacheUpdateErr
	CacheDeleteErr
	CacheFindECache
)

var errMsg = map[uint32]string{
	Successed:         "Successed",
	NotImplementedErr: "Not Implemented Err",

	SendErr:            "Send Err",
	ServeErr:           "ServeErr",
	SafeErr:            "SafeErr",
	InputErr:           "InputErr",
	DataPersistenceErr: "DB Err",
	DomainErr:          "Domain Err",

	InvalidParam:             "Invalid Param",
	MissingParam:             "Missing Param",
	InvalidDelete:            "Invalid Delete",
	InvalidUpdate:            "Invalid InsertArticleVersion",
	ItemNotExist:             "Item not exist",
	UserDuplicateCreationErr: "User Duplicate Creation Err",
	DuplicatePasswordErr:     "Duplicate password",
	EmailFormatErr:           "email format error",
	PermissionDenied:         "Permission Denied",

	JwtSignErr:        "sign tokens err",
	JwtParseErr:       "tokens parse err",
	TokenVerifyErr:    "tokens verify err",
	TokenExpiredErr:   "tokens expired err",
	SuspiciousRequest: "Suspicious Request",

	GenFilenameErr:    "Gen Filename Err",
	MarkdownTOHTMLErr: "Markdown TO HTML Err",
	PwdEncryptionErr:  "password encrytion err",

	EmailSendErr: "Email Send Err",

	DomainEventDataTypeErr: "Domain Event Data Type Err",

	DBCreateErr:     "DB CreateArticle Err",
	DBUpdateErr:     "DB UpdateArticle Err",
	DBDeleteErr:     "DB DeleteArticle Err",
	DBFindErr:       "DB Find Err",
	CacheCreateErr:  "Cache CreateArticle Err",
	CacheUpdateErr:  "Cache UpdateArticle Err",
	CacheDeleteErr:  "Cache DeleteArticle Err",
	CacheFindECache: "Cache Find Err",
}

func init() {
	logger = logging.New("e")
}

type E struct {
	Code uint32
	Prev error
}

func NewError(i uint32, prev error) E {
	return E{
		Code: i,
		Prev: prev,
	}
}

func (e E) Error() string {
	return fmt.Sprintf("%d-%s", e, errMsg[e.Code])
}

func Compare(err error, code uint32) bool {
	var e E
	ok := errors.As(err, &e)
	if !ok {
		return false
	}
	return e.Code == code
}

func ParseErr(err error) (code uint32, message string) {
	var e E
	ok := errors.As(err, &e)
	if !ok {
		return NotImplementedErr, errMsg[NotImplementedErr]
	}
	if e.Code >= 40000 && e.Code <= 99999 {
		logger.Warnf("err:%s", e)
	}
	return e.Code, errMsg[e.Code]
}
