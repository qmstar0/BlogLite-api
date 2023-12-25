package postsCateogry

type UpdatePostCategoryCommand struct {
	Uid           int
	Pid           int
	NewCategoryId int
}

func (u UpdatePostCategoryCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
