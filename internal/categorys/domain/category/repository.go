package category

type CategoryRepository interface {
	Save(category *Category) error
}
