package tags

type UpdateTagCommand struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (u UpdateTagCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
