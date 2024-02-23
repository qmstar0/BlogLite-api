package cqrs

import (
	"context"
	"fmt"
	"strings"
)

type Handler interface {
	Topic() string
	SubscribeTo() any
	Handle(ctx context.Context, v any) error
}

type normHandler[E any] struct {
	topic    string
	handleFn func(ctx context.Context, v *E) error
}

func NewHandler[E any](
	handleFn func(ctx context.Context, v *E) error,
) Handler {
	return &normHandler[E]{
		handleFn: handleFn,
		topic:    getStructName(new(E)),
	}
}

func (n normHandler[E]) Topic() string {
	return n.topic
}
func (n normHandler[E]) Handle(ctx context.Context, v any) error {
	return n.handleFn(ctx, v.(*E))
}
func (n normHandler[E]) SubscribeTo() any {
	return new(E)
}

func getStructName(v any) string {
	t := fmt.Sprintf("%T", v)
	split := strings.Split(t, ".")
	return split[len(split)-1]
}
