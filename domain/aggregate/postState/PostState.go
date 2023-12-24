package postState

import (
	"blog/domain/common"
	"time"
)

type PostState interface {
	Publish(pub common.DomainEventPublisher, cmd PublishPostCommand) error
	Trash(pub common.DomainEventPublisher, cmd TrashPostCommand) error
	UnTrash(pub common.DomainEventPublisher, cmd UnTrashPostCommand) error
}

type PostStateImpl struct {
	Pid       int
	PublishAt int64
	TrashedAt int64
}

func (p *PostStateImpl) Publish(pub common.DomainEventPublisher, cmd PublishPostCommand) error {
	p.PublishAt = time.Now().Unix()
	p.TrashedAt = 0
	return pub.Publish(common.PostPublishedEvent, PostPublishedEvent{
		Uid: cmd.Uid,
		Pid: cmd.Pid,
	})
}

func (p *PostStateImpl) Trash(pub common.DomainEventPublisher, cmd TrashPostCommand) error {
	p.PublishAt = 0
	p.TrashedAt = time.Now().Unix()
	return pub.Publish(common.PostTrashedEvent, PostTrashedEvent{
		Uid: cmd.Uid,
		Pid: cmd.Pid,
	})
}

func (p *PostStateImpl) UnTrash(pub common.DomainEventPublisher, cmd UnTrashPostCommand) error {
	p.PublishAt = time.Now().Unix()
	p.TrashedAt = 0
	return pub.Publish(common.PostRestoredEvent, PostRestoredEvent{
		Uid: cmd.Uid,
		Pid: cmd.Pid,
	})
}
