package event

const (
	DelCache = "<event:DelCache>"
	SendMail = "<event:SendMail>"
)

type DataEvent struct {
	Data  interface{}
	Topic string
	Bus   string
	Time  int64
}
