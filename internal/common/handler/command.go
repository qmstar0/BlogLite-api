package handler

import (
	"context"
)

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
