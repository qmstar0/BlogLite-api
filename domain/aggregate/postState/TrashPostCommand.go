package postState

type TrashPostCommand struct {
	Uid int
	Pid int
}

func (t TrashPostCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
