package entity

type PostTag interface {
	AddTag() error
	RemoveTag() error
}

type PostTagImpl struct {
	Pid string
	Tag []uint
}
