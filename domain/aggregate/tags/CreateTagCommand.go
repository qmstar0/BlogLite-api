package tags

type CreateTagCommand struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (c CreateTagCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
