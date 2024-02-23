package decorator

import "context"

type EventHandler[E any] interface {
	Handle(ctx context.Context, event E) error
}
