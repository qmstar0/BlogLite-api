package postState

type PostStateRepository interface {
	Save(state PostState) error
	Find(pid int) (PostState, error)
}
