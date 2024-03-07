package domainevent

import "time"

type DomainEvent struct {
	AggregateID uint32
	EventID     string
	EventType   uint16
	Event       any
	Timestamp   time.Time
}
