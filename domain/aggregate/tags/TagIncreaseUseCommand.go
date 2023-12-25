package tags

type TagIncreaseUseCommand struct {
	Uid   int
	TagId int
}

func (t TagIncreaseUseCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
