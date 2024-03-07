package adapter

import (
	"common/domainevent"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type CategoryDomainEventStoreModel struct {
	EventID     string
	AggregateID uint32
	Type        uint16
	Event       bson.Raw
	Timestamp   time.Time
}

func (CategoryDomainEventStoreModel) TableName() string {
	return "Domain_EventStore_Cateogry"
}

func eventToModel(event domainevent.DomainEvent) (*CategoryDomainEventStoreModel, error) {
	marshal, err := bson.Marshal(event.Event)
	if err != nil {
		return nil, err
	}
	return &CategoryDomainEventStoreModel{
		EventID:     event.EventID,
		AggregateID: event.AggregateID,
		Type:        event.EventType,
		Event:       marshal,
		Timestamp:   event.Timestamp,
	}, nil
}
