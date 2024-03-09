package adapters

import (
	"common/domainevent"
	"common/e"
	"common/eventstore"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"users/domain/user"
)

const UserCollection = "DevBlogUser"

type userRepositoryImpl struct {
	store    *eventstore.Store
	replayer eventstore.Replayer[user.User]
}

func NewUserRepository(database *mongo.Database) user.UserRepository {
	return &userRepositoryImpl{
		store:    eventstore.NewStore(database, UserCollection),
		replayer: InitUserEventReplayer(),
	}
}

func (r userRepositoryImpl) Find(ctx context.Context, uid uint32) (*user.User, error) {
	events, err := r.store.FindEntityEvents(ctx, uid)
	if err != nil {
		return nil, e.Wrap(e.FindErr, err)
	}

	if len(events) <= 0 {
		return nil, e.Wrap(e.ResourceDoesNotExist, errors.New("resource does not exist"))
	}

	u := new(user.User)
	err = r.replayer.Replay(events, u)
	if err != nil {
		return nil, e.Wrap(e.ReplyEventsErr, err)
	}
	return u, nil
}

func (r userRepositoryImpl) Save(ctx context.Context, u *user.User) error {

	events := user.EventFromAggregate(u)
	if len(events) <= 0 {
		return nil
	}

	err := r.store.StoreEvent(ctx, events)
	if err != nil {
		return e.Wrap(e.StoreEventErr, err)
	}

	if shouldTakeSnapshot(events) {
		if err = r.store.Snapshot(ctx, u.Uid, u); err != nil {
			return e.Wrap(e.SnapshotFailed, err)
		}
	}
	return nil
}

func shouldTakeSnapshot(events []domainevent.DomainEvent) bool {
	for _, event := range events {
		// 检查领域事件是否为资源已创建
		switch event.EventType {
		case domainevent.Created, user.PasswordReset:
			return true
		default:
		}
	}
	return false
}

func InitUserEventReplayer() eventstore.Replayer[user.User] {
	return eventstore.Replayer[user.User]{
		domainevent.Snapshotted: func(raw any, entity *user.User) error {
			event, err := eventstore.Mapping[user.User](raw)
			if err != nil {
				return e.Wrap(e.EventMappingErr, err)
			}
			*entity = *(&event)
			return nil
		},

		//On username updated
		domainevent.Updated: func(raw any, entity *user.User) error {
			event, err := eventstore.Mapping[user.UsernameUpdated](raw)
			if err != nil {
				return e.Wrap(e.EventMappingErr, err)
			}
			if entity.Name.String() != event.OldName {
				return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
			}
			entity.Name = user.UserName(event.NewName)
			return nil
		},

		//On password reset
		user.PasswordReset: func(raw any, entity *user.User) error {
			event, err := eventstore.Mapping[user.UserPasswordReset](raw)
			if err != nil {
				return e.Wrap(e.EventMappingErr, err)
			}
			if entity.Password.String() != event.OldPassowrd {
				return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
			}
			entity.Password = user.Password(event.NewPassword)
			return nil
		},

		//On user rights changed
		user.RightsChanged: func(raw any, entity *user.User) error {
			event, err := eventstore.Mapping[user.UserRightsChanged](raw)
			if err != nil {
				return e.Wrap(e.EventMappingErr, err)
			}
			if uint16(entity.Rights) & ^event.OldRights == 0 {
				return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
			}
			entity.Rights = user.UserRights(event.NewRights)
			return nil
		},
	}
}
