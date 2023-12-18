package event

type CategoryUpdateEvent struct {
	Id          uint
	ParendId    uint
	DisplayName string
	SeoDesc     string
}
