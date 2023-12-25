package categorys

type CreateCategoryCommand struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (c CreateCategoryCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
