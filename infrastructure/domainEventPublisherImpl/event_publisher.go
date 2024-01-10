package domainEventPublisherImpl

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type DomainEventPublisher struct {
	pub *cqrs.EventBus
}

func NewDomainEventPublisher(pub *cqrs.EventBus) *DomainEventPublisher {
	return &DomainEventPublisher{pub: pub}
}

func (d DomainEventPublisher) Publish(event any) error {
	return d.pub.Publish(context.Background(), event)
}
