package event

import (
	"blog/pkg/cqrs"
	"common/events"
	"common/handler"
	"context"
)

type CategoryCreated struct {
	Cid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

type CategoryCreatedHandler handler.EventHandler[*CategoryCreated]

type categoryCreatedHandler struct {
	store events.EventStore
}

func NewCategoryCreatedHandler(store events.EventStore) CategoryCreatedHandler {
	return &categoryCreatedHandler{store: store}
}

func (c categoryCreatedHandler) Handle(ctx context.Context, event *CategoryCreated) error {
	return c.store.Store(ctx, events.Event{
		EventID:     cqrs.GetIdFromCtx(ctx),
		AggregateID: event.Cid,
		Type:        events.DomainEvent,
		Data:        events.Wrap(event),
	})
}
