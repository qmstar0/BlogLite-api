package posts

type PostCreatedEvent struct {
	Uid int
	Pid int
}

func (c PostCreatedEvent) Topic() string {
	return ""
}
