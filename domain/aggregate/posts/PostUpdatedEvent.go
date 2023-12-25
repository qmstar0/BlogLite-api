package posts

type PostUpdatedEvent struct {
	Uid int
	Pid int
}

func (c PostUpdatedEvent) Topic() string {
	return ""
}
