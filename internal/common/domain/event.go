package domain

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

func newEventID() string {
	return uuid.New().String()
}

type DomainEvent struct {
	EventID string
	Payload []byte
	Topic   string
}

func apply(queue *[]*DomainEvent, event any) {
	marshal, _ := json.Marshal(event)
	*queue = append(*queue, &DomainEvent{
		EventID: newEventID(),
		Topic:   EventTopic(event),
		Payload: marshal,
	})
}

func EventTopic(event any) string {
	return fmt.Sprintf("%T", event)
}
