package domain

type BaseAggregate struct {
	events []*DomainEvent
}

func (a *BaseAggregate) Emit(event any) {
	apply(&a.events, event)
}

func (a *BaseAggregate) Events() []*DomainEvent {
	return a.events
}
