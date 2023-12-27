package common

import "context"

type DomainEventPublisher interface {
	Publish(ctx context.Context, event any) error
}
