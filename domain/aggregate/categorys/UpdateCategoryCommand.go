package categorys

type UpdateCategoryCommand struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (u UpdateCategoryCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
