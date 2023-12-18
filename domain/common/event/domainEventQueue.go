package event

type EventQueue interface {
	EnQueue(event Event) error
	Queue() []Event
}
