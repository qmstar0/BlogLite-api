package category

import (
	"common/domainevent"
)

type Category struct {
	Cid         uint32
	Name        Name
	DisplayName string
	SeoDesc     string

	events []domainevent.DomainEvent
}

func CreateCategory(name Name, displayName, seoDesc string) *Category {

	cid := name.ToID()
	return &Category{
		Cid:         cid,
		Name:        name,
		DisplayName: displayName,
		SeoDesc:     seoDesc,
		events: []domainevent.DomainEvent{domainevent.NewDomainEvent(cid, domainevent.Created, &CategoryCreated{
			Name:        name.String(),
			DisplayName: displayName,
			SeoDesc:     seoDesc,
		})},
	}
}

func (c *Category) ChangeCategory(displayName, seoDesc string) {
	if displayName == "" {
		displayName = c.DisplayName
	}
	if seoDesc == "" {
		seoDesc = c.SeoDesc
	}
	//数据完全相同，无需写入数据库
	if c.DisplayName == displayName && c.SeoDesc == seoDesc {
		return
	}
	c.events = append(c.events, domainevent.NewDomainEvent(c.Cid, domainevent.Updated, &CategoryChanged{
		OldDisplayName: c.DisplayName,
		NewDisplayName: displayName,
		OldSeoDesc:     c.SeoDesc,
		NewSeoDesc:     seoDesc,
	}))
	c.SeoDesc = seoDesc
	c.DisplayName = displayName
}

func (c *Category) Delete() {
	name := c.Name.String()
	c.events = append(c.events, domainevent.NewDomainEvent(c.Cid, domainevent.Deleted, &CategoryDeleted{
		Name: name,
	}))
}

func EventFromAggregate(agg *Category) []domainevent.DomainEvent {
	return agg.events
}
