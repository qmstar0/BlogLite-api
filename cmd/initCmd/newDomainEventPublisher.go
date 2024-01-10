package initCmd

import (
	"blog/infrastructure/domainEventPublisherImpl"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

func NewDomainEventPublisher(bus *cqrs.EventBus) *domainEventPublisherImpl.DomainEventPublisher {
	publisher := domainEventPublisherImpl.NewDomainEventPublisher(bus)
	return publisher
}
