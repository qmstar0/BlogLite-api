package common

type DomainEventPublisher interface {
	Publish(event any) error
}
