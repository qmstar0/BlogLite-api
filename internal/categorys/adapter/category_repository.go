package adapter

import (
	"categorys/domain/category"
	"common/domainevent"
	"common/e"
	"common/eventstore"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"users/domain/user"
)

const CategoryCollection = "DevBlogCategory"

type categoryRepositoryImpl struct {
	store    *eventstore.Store
	replayer eventstore.Replayer[category.Category]
}

func NewCategoryRepository(db *mongo.Database) category.CategoryRepository {
	return &categoryRepositoryImpl{
		store:    eventstore.NewStore(db, CategoryCollection),
		replayer: InitCategoryEventReplayer(),
	}
}

func (c categoryRepositoryImpl) Find(ctx context.Context, aggID uint32) (*category.Category, error) {
	events, err := c.store.FindEntityEvents(ctx, aggID)
	if err != nil {
		return nil, e.Wrap(e.FindErr, err)
	}

	if len(events) <= 0 {
		return nil, e.Wrap(e.ResourceDoesNotExist, errors.New("resource does not exist"))
	}

	cate := new(category.Category)
	err = c.replayer.Replay(events, cate)
	if err != nil {
		return nil, e.Wrap(e.ReplyEventsErr, err)
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
		if err := c.store.Snapshot(ctx, cate.Cid, cate); err != nil {
			return e.Wrap(e.SnapshotFailed, err)
		}
	}
	return nil
}

func shouldTakeSnapshot(events []domainevent.DomainEvent) bool {
	for i := range events {
		switch events[i].EventType {
		case domainevent.Created, user.PasswordReset:
			return true
		default:
		}
	}
	return false
}

func InitCategoryEventReplayer() eventstore.Replayer[category.Category] {
	return eventstore.Replayer[category.Category]{

		//On Snapshotted
		domainevent.Snapshotted: func(raw any, entity *category.Category) error {
			event, err := eventstore.Mapping[category.Category](raw)
			if err != nil {
				return e.Wrap(e.EventMappingErr, err)
			}
			*entity = *(&event)
			return nil
		},

		//On Updated
		domainevent.Updated: func(raw any, entity *category.Category) error {
			event, err := eventstore.Mapping[category.CategoryChanged](raw)
			if err != nil {
				return e.Wrap(e.EventMappingErr, err)
			}
			if entity.SeoDesc != event.OldSeoDesc || entity.DisplayName != event.OldDisplayName {
				return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
			}
			entity.SeoDesc, entity.DisplayName = event.NewSeoDesc, event.NewDisplayName
			return nil
		},

		//On Delete
		domainevent.Deleted: func(raw any, entity *category.Category) error {
			return e.Wrap(e.ResourceDoesNotExist, errors.New("no data found"))
		},
	}
}
