package aggregates

import "github.com/qmstar0/domain/internal/domain/domainevent"

type AggregateRoot struct {
	events domainevent.EventQueue
}

func (a AggregateRoot) Emit(event any) {
	domainevent.Emit(a.events, event)
}

func (a AggregateRoot) Events() domainevent.EventQueue {
	return a.events
}
