package eventstore

import (
	"common/domainevent"
	"common/e"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store[E any] struct {
	session    mongo.Session
	collection *mongo.Collection
	replayMap  map[uint16]func(raw any, entity *E) error
}

func NewStore[E any](collection *mongo.Collection) *Store[E] {
	startSession, err := collection.Database().Client().StartSession()
	if err != nil {
		panic(err)
	}

	if err = buildIndexes(collection); err != nil {
		panic(err)
	}

	return &Store[E]{
		session:    startSession,
		collection: collection,
		replayMap:  make(map[uint16]func(raw any, entity *E) error),
	}
}

func (s Store[E]) SetReplayCase(eventType uint16, task func(raw any, entity *E) error) {
	s.replayMap[eventType] = task
}
func (s Store[E]) Replay(events []domainevent.DomainEvent, entity *E) error {
	for _, event := range events {
		if f, ok := s.replayMap[event.EventType]; ok {
			if err := f(event.Event, entity); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s Store[E]) StoreEvent(ctx context.Context, events []domainevent.DomainEvent) error {
	return mongo.WithSession(ctx, s.session, func(sctx mongo.SessionContext) error {
		var err error
		for _, event := range events {
			if event.Event, err = bson.Marshal(event.Event); err != nil {
				return e.Wrap(e.MarshalEventErr, err)
			}
			if _, err = s.collection.InsertOne(ctx, event); err != nil {
				return e.Wrap(e.InsertDataErr, err)
			}
		}
		return nil
	})
}

func (s Store[E]) FindEvent(ctx context.Context, pipeline mongo.Pipeline) ([]domainevent.DomainEvent, error) {
	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, e.Wrap(e.DatabaseErr, err)
	}

	events := make([]domainevent.DomainEvent, 0)
	err = cursor.All(ctx, &events)
	if err != nil {
		return nil, e.Wrap(e.UnmarshalEventErr, err)
	}

	return events, nil
}

func (s Store[E]) Snapshot(ctx context.Context, aggid uint32, etype uint16, entity *E) error {
	return s.StoreEvent(ctx, []domainevent.DomainEvent{domainevent.NewDomainEvent(aggid, etype, entity)})
}

func buildIndexes(collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{"timestamp", -1}, {"aggregateid", 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}
