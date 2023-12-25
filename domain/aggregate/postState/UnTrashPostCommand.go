package postState

type UnTrashPostCommand struct {
	Uid int
	Pid int
}

func (u UnTrashPostCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
