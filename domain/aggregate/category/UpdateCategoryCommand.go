package category

type UpdateCategoryCommand struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}
