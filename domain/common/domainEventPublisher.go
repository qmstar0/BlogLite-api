package common

type DomainEventPublisher interface {
	Publish(topic string, event DomainEvent) error
}
