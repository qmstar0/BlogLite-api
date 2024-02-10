package category

type CategoryRepository interface {
	Save(cate *Category) error
}
