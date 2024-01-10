package categorys

import (
	"blog/domain/common"
	"time"
)

type Category interface {
	Update(pub common.DomainEventPublisher, cmd UpdateCategoryCommand) error
	Delete(pub common.DomainEventPublisher, cmd DeleteCategoryCommand) error
}

type CategoryImpl struct {
	Id          int
	Name        string
	DisplayName string
	SeoDesc     string
	Num         uint
	DeleteAt    int64
}

func CreateCategory(pub common.DomainEventPublisher, cmd CreateCategoryCommand) (Category, error) {
	err := pub.Publish(CategoryCreatedEvent{
		Uid:         cmd.Uid,
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	})
	if err != nil {
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
	return pub.Publish(CategoryUpdatedEvent{
		CategoryId:  c.Id,
		Uid:         cmd.Uid,
		Name:        cmd.Name,
		DisplayName: cmd.DisplayName,
		SeoDesc:     cmd.SeoDesc,
	})
}

func (c *CategoryImpl) Delete(pub common.DomainEventPublisher, cmd DeleteCategoryCommand) error {
	c.DeleteAt = time.Now().Unix()
	return pub.Publish(CategoryDeletedEvent{
		Uid:        cmd.Uid,
		CategoryId: cmd.CategoryId,
	})
}
