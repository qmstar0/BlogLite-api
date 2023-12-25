package postState

type PublishPostCommand struct {
	Uid int
	Pid int
}

func (p PublishPostCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
