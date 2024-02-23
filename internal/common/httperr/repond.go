package httperr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
	InvalidUpdate
	InvalidDelete
	ItemNotExist
	UserDuplicateCreationErr
	DuplicatePasswordErr
	ArticleDuplicateCreationErr
	EmailFormatErr
	PermissionDenied
	LoginRequired
)

// 安全和鉴权错误
const (
	SafeErr StateCode = iota + 31010
	JwtSignErr
	JwtParseErr
	TokenVerifyErr
	TokenExpiredErr
	SuspiciousRequest
)

// 后台错误
const (
	ServeErr StateCode = iota + 40010
	GenFilenameErr
	MarkdownTOHTMLErr
	PwdEncryptionErr
)

// 通信错误
const (
	SendErr StateCode = iota + 51010
	EmailSendErr
)

// 领域服务或事件错误
const (
	DomainErr StateCode = iota + 60010
	DomainEventDataTypeErr
)

// 数据持久化错误
const (
	DataPersistenceErr StateCode = iota + 80010
	DBCreateErr
	DBUpdateErr
	DBDeleteErr
	DBFindErr
	CacheCreateErr
	CacheUpdateErr
	CacheDeleteErr
	CacheFindECache
	ScanSetErr
	ValueGetErr
)

var errMsg = map[StateCode]string{
	Successed: "Successed",

	SendErr:            "Publish Err",
	ServeErr:           "ServeErr",
	SafeErr:            "SafeErr",
	InputErr:           "InputErr",
	DataPersistenceErr: "DB Err",
	DomainErr:          "Domain Err",

	InvalidParam:                "Invalid Param",
	MissingParam:                "Missing Param",
	InvalidDelete:               "Invalid Delete",
	InvalidUpdate:               "Invalid Update",
	ItemNotExist:                "Item not exist",
	UserDuplicateCreationErr:    "User Duplicate Creation Err",
	ArticleDuplicateCreationErr: "Article Duplicate Creation Err",
	DuplicatePasswordErr:        "Duplicate password",
	EmailFormatErr:              "email format httpError",
	PermissionDenied:            "Permission Denied",
	LoginRequired:               "Login Required",

	JwtSignErr:        "sign tokens err",
	JwtParseErr:       "tokens parse err",
	TokenVerifyErr:    "tokens verify err",
	TokenExpiredErr:   "tokens expired err",
	SuspiciousRequest: "Suspicious Request",

	GenFilenameErr:    "Gen Filename Err",
	MarkdownTOHTMLErr: "Markdown TO HTML Err",
	PwdEncryptionErr:  "password encrytion err",

	EmailSendErr: "Email Publish Err",

	DomainEventDataTypeErr: "Domain Event Data Type Err",

	DBCreateErr:     "DB CreateArticle Err",
	DBUpdateErr:     "DB UpdateArticle Err",
	DBDeleteErr:     "DB DeleteArticle Err",
	DBFindErr:       "DB Find Err",
	CacheCreateErr:  "Cache CreateArticle Err",
	CacheUpdateErr:  "Cache UpdateArticle Err",
	CacheDeleteErr:  "Cache DeleteArticle Err",
	CacheFindECache: "Cache Find Err",
	ScanSetErr:      "Set Scan Err",
	ValueGetErr:     "Value Get Err",
}

func (s StateCode) Error() string {
	msg, _ := errMsg[s]
	return msg
}
func Error(code StateCode, info string) error {
	return fmt.Errorf("%w (%s)", code, info)
}

type response struct {
	Code    StateCode
	Message string
}

func Respond(w http.ResponseWriter, e *error) {
	resp := response{
		Code:    Successed,
		Message: Successed.Error(),
	}
	if *e != nil {
		if errors.As(errors.Unwrap(*e), &resp.Code) {
			resp.Message = resp.Code.Error()
		} else {
			resp.Code = NotImplementedErr
			resp.Message = (*e).Error()
		}
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
