package bus

import "context"

type bus[T any] interface {
	Publish(ctx context.Context, v T) error
}

type CommandBus bus[any]
type QueryBus bus[any]
