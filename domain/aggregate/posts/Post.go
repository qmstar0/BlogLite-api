package posts

import (
	"blog/domain/common"
	"time"
)

type Post interface {
	Update(pub common.DomainEventPublisher, cmd UpdatePostCommand) error
	Delete(pub common.DomainEventPublisher, cmd DeletePostCommand) error
}

type PostImpl struct {
	Pid      int
	Title    string
	Slug     string
	Summary  string
	Original string
	Content  string
	DeleteAt int64
}

func NewPostDraft(pub common.DomainEventPublisher, cmd NewPostDraftCommand) (Post, error) {
	pid := GenPId.NextID()
	if err := pub.Publish(PostCreatedEvent{
		Uid: cmd.Uid,
		Pid: pid,
	}); err != nil {
		return nil, err
	}
	return &PostImpl{
		Pid: pid,
	}, nil
}

func (p *PostImpl) Update(pub common.DomainEventPublisher, cmd UpdatePostCommand) error {
	p.Title = cmd.Title
	p.Original = cmd.Content
	p.Slug = cmd.Slug
	p.Content = ""
	return pub.Publish(PostUpdatedEvent{
		Uid: cmd.Uid,
		Pid: p.Pid,
	})
}

func (p *PostImpl) Delete(pub common.DomainEventPublisher, cmd DeletePostCommand) error {
	p.DeleteAt = time.Now().Unix()
	return pub.Publish(PostDeletedEvent{
		Uid: cmd.Uid,
		Pid: p.Pid,
	})
}
