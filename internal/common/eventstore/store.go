package eventstore

import (
	"common/domainevent"
	"common/e"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	session    mongo.Session
	collection *mongo.Collection
}

func NewStore(database *mongo.Database, collectionName string) *Store {
	session, err := database.Client().StartSession()
	if err != nil {
		panic(err)
	}

	collection := database.Collection(collectionName)

	if err = buildIndexes(collection); err != nil {
		panic(err)
	}

	return &Store{
		session:    session,
		collection: collection,
	}
}
func (s Store) FindEntityEvents(ctx context.Context, aggid uint32) ([]domainevent.DomainEvent, error) {
	cursor, err := s.collection.Aggregate(ctx, DefaultEntityQuery(aggid))
	if err != nil {
		return nil, e.Wrap(e.FindErr, err)
	}
	var events = make([]domainevent.DomainEvent, 0)

	err = cursor.All(ctx, &events)
	if err != nil {
		return nil, e.Wrap(e.FindResultToModelsErr, err)
	}
	return events, nil
}

func (s Store) StoreEvent(ctx context.Context, events []domainevent.DomainEvent) error {
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

func (s Store) Snapshot(ctx context.Context, aggid uint32, entity any) error {
	return s.StoreEvent(ctx, []domainevent.DomainEvent{domainevent.NewDomainEvent(aggid, domainevent.Snapshotted, entity)})
}

func (s Store) AutoSnapshot(ctx context.Context) {

}

//func (s Store) Session(ctx context.Context, task SessionTask) error {
//	return mongo.WithSession(ctx, s.session, func(sctx mongo.SessionContext) error {
//		return task(sctx, s.collection)
//	})
//}
//
//func (s Store) Collection() *mongo.Collection {
//	return s.collection
//}

type SessionTask func(sctx mongo.SessionContext, coll *mongo.Collection) error
type CommandTask func(ctx context.Context, coll *mongo.Collection) error

func buildIndexes(collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{"timestamp", -1}, {"aggregateid", 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}
