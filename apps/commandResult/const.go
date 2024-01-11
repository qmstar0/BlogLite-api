package commandResult

const Successed StateCode = 0

// NotImplementedErr 临时占位错误
const NotImplementedErr StateCode = 99999

// CQRS相关
const (
	EventPublishErr StateCode = iota + 10010
	EventProcessErr
	CommandPuiblishErr
	CommandProcessErr
	WaitReplyTimeoutErr
)

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
	DBSaveErr
	DBDeleteErr
	DBFindErr
	CacheCreateErr
	CacheUpdateErr
	CacheDeleteErr
	CacheFindECache
	ScanSetErr
	ValueGetErr
)
