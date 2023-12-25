package postsTags

type UpdatePostTagsCommand struct {
	Uid    int
	Pid    int
	TagIds []int
}

func (u UpdatePostTagsCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
