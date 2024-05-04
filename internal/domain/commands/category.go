package commands

type CreateCategory struct {
	Name string `cli:"textinput"`
	Desc string `cli:"textinput"`
}

type ModifyCategoryDesc struct {
	ID      uint32 `cli:"select"`
	NewDesc string `cli:"textinput"`
}
type DeleteCategory struct {
	ID uint32 `cli:"select"`
}
