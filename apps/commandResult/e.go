package commandResult

var MessageMap = map[StateCode]string{
	Successed:         "Successed",
	NotImplementedErr: "Not Implemented Err",

	SendErr:            "Send Err",
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

	EmailSendErr: "Email Send Err",

	DomainEventDataTypeErr: "Domain Event Data Type Err",

	DBDeleteErr:     "DB DeleteArticle Err",
	DBFindErr:       "DB Find Err",
	CacheCreateErr:  "Cache CreateArticle Err",
	CacheUpdateErr:  "Cache UpdateArticle Err",
	CacheDeleteErr:  "Cache DeleteArticle Err",
	CacheFindECache: "Cache Find Err",
	ScanSetErr:      "Set Scan Err",
	ValueGetErr:     "Value Get Err",
}
