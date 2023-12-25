package postsTags

import "blog/domain/common"

type PostTags interface {
	Update(pub common.DomainEventPublisher, cmd UpdatePostTagsCommand) error
}

type PostTagsImpl struct {
	Pid    int
	TagIds []int
}

func (p *PostTagsImpl) Update(pub common.DomainEventPublisher, cmd UpdatePostTagsCommand) error {
	p.TagIds = cmd.TagIds
	return pub.Publish(PostTagsUpdated{
		Uid:       cmd.Uid,
		Pid:       p.Pid,
		OldTagIds: p.TagIds,
		NewTagIds: cmd.TagIds,
	})
}
