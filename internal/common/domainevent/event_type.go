package domainevent

const (
	Snapshotted uint16 = 1 << iota

	Created
	Updated
	Deleted
)
