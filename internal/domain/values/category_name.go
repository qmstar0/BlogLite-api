package values

type CategoryName string

func NewCategoryName(s string) (CategoryName, error) {
	return CategoryName(s), nil
}

func (n CategoryName) String() string {
	return string(n)
}
