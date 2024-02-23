package adapter

import (
	"common/events"
	"time"
)

type CategoryDomainEventStoreModel struct {
	Id          int
	EventID     string
	AggregateID int
	Type        int16
	Data        []byte
	timestamp   time.Time
}

func eventToModel(event events.Event) (*CategoryDomainEventStoreModel, error) {
	data, err := event.Data.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return &CategoryDomainEventStoreModel{
		EventID:     event.EventID,
		AggregateID: event.AggregateID,
		Type:        int16(event.Type),
		Data:        data,
		timestamp:   time.Time{},
	}, nil
}
