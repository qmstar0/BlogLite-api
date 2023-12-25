package postsCateogry

import "blog/domain/common"

type PostCategory interface {
	Remove(pub common.DomainEventPublisher, cmd RemovePostCategoryCommand) error
	Update(pub common.DomainEventPublisher, cmd UpdatePostCategoryCommand) error
}

type PostCategoryImpl struct {
	Pid        int
	CategoryId int
}

func (p *PostCategoryImpl) Remove(pub common.DomainEventPublisher, cmd RemovePostCategoryCommand) error {
	p.CategoryId = 0
	return pub.Publish(PostCategoryRemovedEvent{
		Uid:           cmd.Uid,
		Pid:           p.Pid,
		OldCategoryId: p.CategoryId,
	})
}

func (p *PostCategoryImpl) Update(pub common.DomainEventPublisher, cmd UpdatePostCategoryCommand) error {
	p.CategoryId = cmd.NewCategoryId
	return pub.Publish(PostCategoryUpdatedEvent{
		Uid:           cmd.Uid,
		Pid:           p.Pid,
		OldCategoryId: p.CategoryId,
		NewCategoryID: cmd.NewCategoryId,
	})
}
