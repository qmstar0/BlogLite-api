package cqrs

import "context"

type Publisher interface {
	Publish(c context.Context, v any) error
}
