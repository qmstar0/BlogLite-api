package category

type CategoryUpdatedEvent struct {
	Uid         int
	CategoryId  int
	Name        string
	DisplayName string
	SeoDesc     string
}
