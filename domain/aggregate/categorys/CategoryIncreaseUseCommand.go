package categorys

type CategoryIncreaseUseCommand struct {
	Uid        int
	CategoryId int
}

func (c CategoryIncreaseUseCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
