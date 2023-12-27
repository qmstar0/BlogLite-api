package tags

type TagUpdatedEvent struct {
	Uid         int
	TagId       int
	Name        string
	DisplayName string
	SeoDesc     string
}
