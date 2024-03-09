package category

type CategoryCreated struct {
	Name        string
	DisplayName string
	SeoDesc     string
}

type CategoryChanged struct {
	OldDisplayName string
	NewDisplayName string
	OldSeoDesc     string
	NewSeoDesc     string
}

type CategoryDeleted struct {
	Name string
}
