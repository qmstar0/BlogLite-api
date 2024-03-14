package adapter

import (
	"blog/pkg/mongodb"
	"blog/pkg/rediscache"
	"blog/pkg/shutdown"
	"categorys/domain/category"
	"common/domainevent"
	"common/e"
	"common/eventstore"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"users/domain/user"
)

func init() {
	dbCloseFn := mongodb.Init()
	shutdown.OnShutdown(func() error { return dbCloseFn(context.Background()) })

	cacheCloseFn := rediscache.Init()
	shutdown.OnShutdown(cacheCloseFn)
}

const CategoryCollection = "DevBlogCategory"

type categoryRepositoryImpl struct {
	store *eventstore.Store[category.Category]
}

func NewCategoryRepository(db *mongo.Database) category.CategoryRepository {
	collection := db.Collection(CategoryCollection)
	store := eventstore.NewStore[category.Category](collection)
	InitCategoryReplayer(store)
	return &categoryRepositoryImpl{
		store: store,
	}
}

func (c categoryRepositoryImpl) Exist(ctx context.Context, aggID uint32) (bool, error) {
	events, err := c.store.FindEvent(ctx, checkCategoryIsExistQ(aggID))
	if err != nil {
		return false, e.Wrap(e.FindEventErr, err)
	}
	return len(events) > 0, nil
}

func (c categoryRepositoryImpl) Find(ctx context.Context, aggID uint32) (*category.Category, error) {
	events, err := c.store.FindEvent(ctx, getCategoryQ(aggID))
	if err != nil {
		return nil, e.Wrap(e.FindEventErr, err)
	}

	if len(events) <= 0 {
		return nil, e.Wrap(e.ResourceDoesNotExist, e.FindResultIsNull)
	}

	cate := new(category.Category)
	err = c.store.Replay(events, cate)
	if err != nil {
		return nil, e.Wrap(e.ReplayEventsErr, err)
	}
	return cate, nil
}

func (c categoryRepositoryImpl) Save(ctx context.Context, cate *category.Category) error {
	events := category.EventFromAggregate(cate)
	if len(events) <= 0 {
		return nil
	}

	// 普通事件持久化和快照事件持久化不放在一个事务里
	// 是因为在合理情况下，快照允许失败，但快照失败不应该影响普通事件的持久化
	if err := c.store.StoreEvent(ctx, events); err != nil {
		return e.Wrap(e.StoreEventErr, err)
	}

	if shouldTakeSnapshot(events) {
		if err := c.store.Snapshot(ctx, cate.Cid, category.Snapshotted, cate); err != nil {
			return e.Wrap(e.SnapshotFailed, err)
		}
	}
	return nil
}

func shouldTakeSnapshot(events []domainevent.DomainEvent) bool {
	for i := range events {
		switch events[i].EventType {
		case category.Created, user.PasswordReset:
			return true
		default:
		}
	}
	return false
}

func InitCategoryReplayer(store *eventstore.Store[category.Category]) {

	//category.Category
	store.SetReplayCase(category.Snapshotted, func(raw any, entity *category.Category) error {
		event, err := eventstore.Mapping[category.Category](raw)
		if err != nil {
			return e.Wrap(e.EventMappingErr, err)
		}
		*entity = *(&event)
		return nil
	})

	//category.CategoryChanged
	store.SetReplayCase(category.Updated, func(raw any, entity *category.Category) error {
		event, err := eventstore.Mapping[category.CategoryChanged](raw)
		if err != nil {
			return e.Wrap(e.EventMappingErr, err)
		}
		if entity.SeoDesc != event.OldSeoDesc || entity.DisplayName != event.OldDisplayName {
			return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
		}
		entity.SeoDesc, entity.DisplayName = event.NewSeoDesc, event.NewDisplayName
		return nil
	})

	//delete
	store.SetReplayCase(category.Deleted, func(raw any, entity *category.Category) error {
		return e.Wrap(e.ResourceDoesNotExist, e.FindResultIsNull)
	})
}
