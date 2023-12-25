package common

type DomainEventPublisher interface {
	Publish(event DomainEvent) error
}
