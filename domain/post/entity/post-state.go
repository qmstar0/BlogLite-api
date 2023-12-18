package entity

import "time"

type PostState interface {
	Publish()
	Trash()
	UnTrash()
}

type PostStateImpl struct {
	Pid       string
	PublishAt int64
	DeleteAt  int64
}

func (p *PostStateImpl) Publish() {
	p.PublishAt = time.Now().Unix()
	p.DeleteAt = 0
}

func (p *PostStateImpl) Trash() {
	p.PublishAt = 0
	p.DeleteAt = time.Now().Unix()
}

func (p *PostStateImpl) UnTrash() {
	p.Publish()
}
