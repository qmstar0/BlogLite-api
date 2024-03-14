package post

const (
	DraftState State = iota
	PublishedState
	TrashState
)

type State uint16

func (s State) IsDraft() bool {
	return s == DraftState
}

func (s State) IsPublished() bool {
	return s == PublishedState
}

func (s State) IsTrash() bool {
	return s == TrashState
}
