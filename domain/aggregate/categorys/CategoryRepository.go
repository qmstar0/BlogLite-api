package categorys

type CategoryRepository interface {
	Save(cate Category) error
	FindById(id int) (Category, error)
	FindByName(name string) (Category, error)
}
