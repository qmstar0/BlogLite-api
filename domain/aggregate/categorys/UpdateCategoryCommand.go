package categorys

type UpdateCategoryCommand struct {
	Uid         int
	CategoryId  int
	Name        string
	DisplayName string
	SeoDesc     string
}
