package postsCateogry

type RemovePostCategoryCommand struct {
	Uid int
	Pid int
}

func (r RemovePostCategoryCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
