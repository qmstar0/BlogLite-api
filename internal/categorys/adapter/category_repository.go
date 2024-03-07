package adapter

import (
	"categorys/domain/category"
	"common/domainevent"
	"common/e"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CategoryRepositoryImpl struct {
	db      *mongo.Collection
	session mongo.Session
}

func NewCategoryRepository(db *mongo.Database) category.CategoryRepository {
	session, err := db.Client().StartSession()
	if err != nil {
		panic(err)
	}

	c := &CategoryRepositoryImpl{
		db:      db.Collection(CategoryDomainEventStoreModel{}.TableName()),
		session: session,
	}
	c.Init()

	return c
}
func (c CategoryRepositoryImpl) Init() {

	initAggregateQuery()
	err := c.BuildIndex(context.Background())
	if err != nil {
		panic(err)
	}
}

func (c CategoryRepositoryImpl) Find(ctx context.Context, aggID uint32) (*category.Category, error) {
	cursor, err := c.db.Aggregate(ctx, queryToGetLastSnapshotAndLastEvent(aggID))
	if err != nil {
		return nil, e.Wrap(e.FindErr, err)
	}

	var models = make([]*CategoryDomainEventStoreModel, 0)

	err = cursor.All(ctx, &models)
	if err != nil {
		return nil, e.Wrap(e.FindResultToModelsErr, err)
	}

	if len(models) <= 0 {
		return nil, e.Wrap(e.FindResultIsNull, errors.New("no data found"))
	}

	cate := new(category.Category)
	err = replyEvents(models, cate)
	if err != nil {
		return nil, e.Wrap(e.ReplyEventsErr, err)
	}
	return cate, nil
}

func (c CategoryRepositoryImpl) Save(ctx context.Context, cate *category.Category) error {
	event := category.EventFromAggregate(cate)
	if len(event) <= 0 {
		return nil
	}

	models := make([]*CategoryDomainEventStoreModel, len(event))
	for i := range event {
		model, err := eventToModel(event[i])
		if err != nil {
			return e.Wrap(e.EventToModelErr, err)
		}
		models[i] = model
	}

	// 普通事件持久化和快照事件持久化不放在一个事务里
	// 是因为在合理情况下，快照允许失败，但快照失败不应该影响普通事件的持久化
	err := c.StoreEvent(ctx, models)
	if err != nil {
		return e.Wrap(e.StoreEventErr, err)
	}

	if shouldTakeSnapshot(event) {
		err = c.Snapshot(ctx, cate)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c CategoryRepositoryImpl) StoreEvent(ctx context.Context, models []*CategoryDomainEventStoreModel) error {
	if len(models) <= 0 {
		return nil
	}
	err := mongo.WithSession(ctx, c.session, func(sctx mongo.SessionContext) error {
		for i := range models {
			if _, err := c.db.InsertOne(ctx, models[i]); err != nil {
				return e.Wrap(e.InsertDataErr, err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("StoreEvent err", err)
		return err
	}
	return nil
}

func (c CategoryRepositoryImpl) Snapshot(ctx context.Context, cate *category.Category) error {
	cid := cate.Cid
	marshal, err := bson.Marshal(category.CategorySnapshot{
		Cid:         cid,
		Name:        cate.Name.String(),
		DisplayName: cate.DisplayName,
		SeoDesc:     cate.SeoDesc,
	})
	if err != nil {
		return e.Wrap(e.MarshalSnapshotEventErr, err)
	}

	models := append([]*CategoryDomainEventStoreModel(nil), &CategoryDomainEventStoreModel{
		EventID:     domainevent.NewEventID(),
		AggregateID: cid,
		Type:        domainevent.Snapshotted,
		Event:       marshal,
		Timestamp:   time.Now(),
	})

	return c.StoreEvent(ctx, models)
}

func (c CategoryRepositoryImpl) BuildIndex(ctx context.Context) error {
	_, err := c.db.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"timestamp", -1}, {"aggregateid", 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func shouldTakeSnapshot(models []domainevent.DomainEvent) bool {
	for i := range models {
		// 检查领域事件是否为资源已创建
		if models[i].EventType&domainevent.Created == domainevent.Created {
			return true
		}
	}
	return false
}

func replyEvents(models []*CategoryDomainEventStoreModel, cate *category.Category) error {
	for i := range models {
		err := loadEventToAggregate(models[i], cate)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadEventToAggregate(model *CategoryDomainEventStoreModel, cate *category.Category) error {
	switch model.Type {

	case domainevent.Snapshotted:
		var event category.CategorySnapshot
		err := bson.Unmarshal(model.Event, &event)
		if err != nil {
			return err
		}
		cate.SeoDesc = event.SeoDesc
		cate.Name = category.Name(event.Name)
		cate.DisplayName = event.DisplayName
		cate.Cid = event.Cid

	case domainevent.Updated:
		var event category.CategoryChanged
		err := bson.Unmarshal(model.Event, &event)
		if err != nil {
			return err
		}

		if cate.SeoDesc != event.OldSeoDesc || cate.DisplayName != event.OldDisplayName {
			return fmt.Errorf("在事件回溯过程中出现了数据错乱")
		}

		cate.SeoDesc, cate.DisplayName = event.NewSeoDesc, event.NewDisplayName

	case domainevent.Deleted:
		//var event category.CategoryDeleted
	case domainevent.Created:
		//var event category.CategoryCreated
	default:
		return fmt.Errorf("未知的事件类型:%d", model.Type)
	}
	return nil
}
