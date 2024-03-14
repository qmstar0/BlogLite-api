package category

import "common/domainevent"

type Category struct {
	ID      uint32
	Name    Name
	SeoDesc string

	events []domainevent.DomainEvent
}

func NewCategory(name Name, seoDesc string) *Category {
	id := name.ToID()
	return &Category{
		ID:      id,
		Name:    name,
		SeoDesc: seoDesc,
		events: []domainevent.DomainEvent{domainevent.NewDomainEvent(id, Created, CategoryCretaed{
			Name:    name.String(),
			SeoDesc: seoDesc,
		})},
	}
}

func (c *Category) ModifySeoDesc(newSeoDesc string) {
	c.events = append(c.events, domainevent.NewDomainEvent(c.ID, SeoDescChanged, CategorySeoDescChanged{
		ID:  c.ID,
		Old: c.SeoDesc,
		New: newSeoDesc,
	}))
	c.SeoDesc = newSeoDesc
}

func (c *Category) Delete() {
	c.events = append(c.events, domainevent.NewDomainEvent(c.ID, Deleted, CategoryDeleted{
		ID: c.ID,
	}))
}
func EventsFromEntity(entity *Category) []domainevent.DomainEvent {
	return entity.events
}
