package entity

type PostCategory interface {
	UpdateCategory(id uint) error
}

type PostCategoryImpl struct {
	Pid      string
	Category uint
}
