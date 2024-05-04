package domainevent

import (
	"context"
	"github.com/google/uuid"
	"time"
)

func NewEventID() string {
	return uuid.New().String()
}

type DomainEvent interface {
	Timestamp() time.Time
	EventID() string
	Payload() any
}

type domainEvent struct {
	timestamp time.Time
	eventID   string
	payload   any
}

func (d *domainEvent) Timestamp() time.Time {
	return d.timestamp
}

func (d *domainEvent) EventID() string {
	return d.eventID
}

func (d *domainEvent) Payload() any {
	return d.payload
}

func Emit(queue EventQueue, event any) {
	queue = append(queue, &domainEvent{
		timestamp: time.Now(),
		eventID:   NewEventID(),
		payload:   event,
	})
}

type EventQueue []DomainEvent

type DomainEventListener interface {
	OnEvent(ctx context.Context, event DomainEvent) error
}

type DomainEventDispatcher interface {
	Dispatch(ctx context.Context, queue EventQueue)
}
