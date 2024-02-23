package events

type EventType uint16

const (
	DomainEvent EventType = iota
	SystemEvent
	MonitoringEvent
	TimingEvent
)
