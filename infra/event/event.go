package event

const (
	DelCache = "<commonEvent:DelCache>"
	SendMail = "<commonEvent:SendMail>"
)

type DataEvent struct {
	Data  interface{}
	Topic string
	Bus   string
	Time  int64
}
