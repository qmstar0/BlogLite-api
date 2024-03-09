package e

type StateCode string

func (e StateCode) Error() string {
	s, _ := errMap[e]
	return s
}

const Successed StateCode = "OK"
const NotImplementedErr StateCode = "?"

// prots 协议层
const (
	InvalidParam      StateCode = "P0101"
	Unauthortion      StateCode = "P0201"
	LoginRequired     StateCode = "P0301"
	LoginExpired      StateCode = "P0302"
	CommandHandlerErr StateCode = "P1001"
	QueryHandlerErr   StateCode = "P2001"

	NewValueObjectErr    StateCode = "D3001"
	ValueObjectCheckErr  StateCode = "D3002"
	ResourceDoesNotExist StateCode = "D1001"
	ResourceCreated      StateCode = "D1002"

	PasswordFormatErr StateCode = "DU1001"
)

const (
	DatabaseErr           StateCode = "S1000"
	FindResultToModelsErr StateCode = "S0901"
	ReplyEventsErr        StateCode = "S0902"
	EventMappingErr       StateCode = "S0903"
	MarshalEventErr       StateCode = "S0904"
	EventDisorder         StateCode = "S0905"
	FindResultIsNull      StateCode = "S0101"
	FindErr               StateCode = "S0102"
	StoreEventErr         StateCode = "S0103"
	InsertDataErr         StateCode = "S0104"
	FindEntityErr         StateCode = "S0105"
	SnapshotFailed        StateCode = "S0106"
)

var errMap = map[StateCode]string{
	InvalidParam:         "Invalid Param",
	Unauthortion:         "Unauthortion",
	CommandHandlerErr:    "err on processing cmd",
	QueryHandlerErr:      "err on processing query",
	NewValueObjectErr:    "err on check data format",
	ResourceDoesNotExist: "resource does not exist",
	ResourceCreated:      "resource is created",
	ValueObjectCheckErr:  "data format error",
	LoginExpired:         "login has expired",
	LoginRequired:        "login required",

	//system err
	EventDisorder:         "error",
	DatabaseErr:           "error",
	FindResultToModelsErr: "error",
	ReplyEventsErr:        "error",
	EventMappingErr:       "error",
	MarshalEventErr:       "error",
	FindErr:               "error",
	FindEntityErr:         "error",
	StoreEventErr:         "error",
	InsertDataErr:         "error",
	SnapshotFailed:        "error",
}

func GetStateCodeMap() map[StateCode]string {
	return errMap
}
