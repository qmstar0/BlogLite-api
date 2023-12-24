package category

import (
	"blog/domain/common"
	"time"
)

type Category interface {
	Update(pub common.DomainEventPublisher, cmd UpdateCategoryCommand) error
	Delete(pub common.DomainEventPublisher, cmd DeleteCategoryCommand) error
	IncreaseUsed(pub common.DomainEventPublisher, cmd CategoryIncreaseUseCommand) error
	ReduceUsed(pub common.DomainEventPublisher, cmd CategoryReduceUseCommand) error
}

type CategoryImpl struct {
	Id          int
	Name        string
	DisplayName string
	SeoDesc     string
	DeleteAt    int64
	Num         int
}

func CreateCategory(pub common.DomainEventPublisher, cmd CreateCategoryCommand) (Category, error) {
	if err := pub.Publish(common.CategoryCreatedEvent, CategoryCreatedEvent{
		Uid:         cmd.Uid,
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	}); err != nil {
		return nil, err
	}
	return &CategoryImpl{
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	}, nil
}
func (c *CategoryImpl) Update(pub common.DomainEventPublisher, cmd UpdateCategoryCommand) error {
	c.Name = cmd.Name
	c.DisplayName = cmd.DisplayName
	c.SeoDesc = cmd.SeoDesc
	return pub.Publish(common.CategoryUpdatedEvent, CategoryUpdatedEvent{
		CategoryId:  c.Id,
		Uid:         cmd.Uid,
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	})
}

func (c *CategoryImpl) Delete(pub common.DomainEventPublisher, cmd DeleteCategoryCommand) error {
	c.DeleteAt = time.Now().Unix()
	return pub.Publish(common.CategoryDeletedEvent, CategoryDeletedEvent{
		Uid:        cmd.Uid,
		CategoryId: cmd.CategoryId,
	})
}

func (c *CategoryImpl) IncreaseUsed(pub common.DomainEventPublisher, cmd CategoryIncreaseUseCommand) error {
	c.Num += 1
	return pub.Publish(common.CategoryIncreasedUsageEvent, CategoryIncreasedUsageEvent{
		Uid:        cmd.Uid,
		CategoryId: cmd.CategoryId,
	})
}

func (c *CategoryImpl) ReduceUsed(pub common.DomainEventPublisher, cmd CategoryReduceUseCommand) error {
	c.Num -= 1
	return pub.Publish(common.CategoryReducedUsageEvent, CategoryReducedUsageEvent{
		Uid:        cmd.Uid,
		CategoryId: cmd.CateogryId,
	})
}
