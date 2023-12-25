package tags

type DeleteTagCommand struct {
	Uid   int
	TagId int
}

func (d DeleteTagCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
