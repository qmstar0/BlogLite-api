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
	LoginExpired      StateCode = "P0301"
	CommandHandlerErr StateCode = "P1001"
	QueryHandlerErr   StateCode = "P2001"

	NewValueObjectErr       StateCode = "D3001"
	ValueObjectCheckErr     StateCode = "D3002"
	ResourceDoesNotExist    StateCode = "D1001"
	ResourceCreated         StateCode = "D1002"
	DatabaseErr             StateCode = "S1000"
	FindResultToModelsErr   StateCode = "S0901"
	ReplyEventsErr          StateCode = "S0902"
	EventToModelErr         StateCode = "S0903"
	MarshalSnapshotEventErr StateCode = "S0904"
	FindResultIsNull        StateCode = "S0101"
	FindErr                 StateCode = "S0102"
	StoreEventErr           StateCode = "S0103"
	InsertDataErr           StateCode = "S0104"
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
}
