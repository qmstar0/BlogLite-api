package tag

type TagRepository interface {
	Save(tag Tag) error
	FindById(id int) (Tag, error)
	FindByName(name string) (Tag, error)
}
