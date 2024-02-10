package bus

import "context"

type CommandBus interface {
	Publish(ctx context.Context, cmd any) error
}
type QueryBus interface {
	Publish(ctx context.Context, query any) error
}
