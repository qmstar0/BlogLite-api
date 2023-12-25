package posts

type DeletePostCommand struct {
	Uid int
	Pid int
}

func (d DeletePostCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
