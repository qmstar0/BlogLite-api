package domainevent

import (
	"common/idtools"
	"time"
)

type DomainEvent struct {
	AggregateID uint32
	EventID     string
	EventType   uint16
	Event       any
	Timestamp   time.Time
}

func NewDomainEvent(aggid uint32, etype uint16, event any) DomainEvent {
	return DomainEvent{
		AggregateID: aggid,
		EventID:     idtools.NewUUID(),
		EventType:   etype,
		Event:       event,
		Timestamp:   time.Now(),
	}
}
