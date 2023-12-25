package tags

import (
	"blog/domain/common"
	"time"
)

type Tag interface {
	Update(pub common.DomainEventPublisher, cmd UpdateTagCommand) error
	Delete(pub common.DomainEventPublisher, cmd DeleteTagCommand) error
	IncreaseUsed(pub common.DomainEventPublisher, cmd TagIncreaseUseCommand) error
	ReduceUsed(pub common.DomainEventPublisher, cmd TagReduceUseCommand) error
}

type TagImpl struct {
	Id          int
	Name        string
	DisplayName string
	SeoDesc     string
	DeleteAt    int64
	Num         int
}

func NewTag(pub common.DomainEventPublisher, cmd CreateTagCommand) (Tag, error) {
	if err := pub.Publish(TagCreatedEvent{
		Uid:         cmd.Uid,
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	}); err != nil {
		return nil, err
	}
	return &TagImpl{
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	}, nil
}

func (t *TagImpl) Update(pub common.DomainEventPublisher, cmd UpdateTagCommand) error {
	t.DisplayName = cmd.DisplayName
	t.SeoDesc = cmd.SeoDesc
	t.Name = cmd.Name
	return pub.Publish(TagUpdatedEvent{
		Uid:         cmd.Uid,
		TagId:       t.Id,
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	})
}

func (t *TagImpl) Delete(pub common.DomainEventPublisher, cmd DeleteTagCommand) error {
	t.DeleteAt = time.Now().Unix()
	return pub.Publish(TagDeletedEvent{
		Uid:   cmd.Uid,
		TagId: t.Id,
	})
}

func (t *TagImpl) IncreaseUsed(pub common.DomainEventPublisher, cmd TagIncreaseUseCommand) error {
	t.Num += 1
	return pub.Publish(TagIncreasedUsageEvent{
		Uid:   cmd.Uid,
		TagId: t.Id,
	})
}

func (t *TagImpl) ReduceUsed(pub common.DomainEventPublisher, cmd TagReduceUseCommand) error {
	t.Num -= 1
	return pub.Publish(TagReducedUsageEvent{
		Uid:   cmd.Uid,
		TagId: t.Id,
	})
}
