package categorys

type DeleteCategoryCommand struct {
	Uid        int
	CategoryId int
}

func (d DeleteCategoryCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
