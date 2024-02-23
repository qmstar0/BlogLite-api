package events

import (
	"context"
)

type EventStore interface {
	Store(ctx context.Context, event Event) error
}
