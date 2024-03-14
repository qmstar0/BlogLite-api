package adapters

import (
	"blog/pkg/mongodb"
	"blog/pkg/rediscache"
	"blog/pkg/shutdown"
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

const UserCollection = "DevBlogUser"

type userRepositoryImpl struct {
	store *eventstore.Store[user.User]
}

func NewUserRepository(database *mongo.Database) user.UserRepository {
	collection := database.Collection(UserCollection)
	store := eventstore.NewStore[user.User](collection)
	//InitCategoryReplayer(store)
	return &userRepositoryImpl{
		store: store,
	}
}

func (r userRepositoryImpl) Exist(ctx context.Context, uid uint32) (bool, error) {
	events, err := r.store.FindEvent(ctx, checkUserIsExistQ(uid))
	if err != nil {
		return false, e.Wrap(e.FindEventErr, err)
	}
	return len(events) > 0, nil
}

func (r userRepositoryImpl) Find(ctx context.Context, uid uint32) (*user.User, error) {
	events, err := r.store.FindEvent(ctx, getUserQ(uid))
	if err != nil {
		return nil, e.Wrap(e.FindEventErr, err)
	}

	if len(events) <= 0 {
		return nil, e.Wrap(e.ResourceDoesNotExist, e.FindResultIsNull)
	}

	u := new(user.User)
	err = r.store.Replay(events, u)
	if err != nil {
		return nil, e.Wrap(e.ReplayEventsErr, err)
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
		if err = r.store.Snapshot(ctx, u.Uid, user.Login, u); err != nil {
			return e.Wrap(e.SnapshotFailed, err)
		}
	}
	return nil
}

func shouldTakeSnapshot(events []domainevent.DomainEvent) bool {
	for _, event := range events {
		switch event.EventType {
		case user.RegistrationSuccess, user.PasswordReset:
			return true
		default:
		}
	}
	return false
}

func InitUserEventReplayer(store eventstore.Store[user.User]) {
	store.SetReplayCase(user.Login, func(raw any, entity *user.User) error {
		event, err := eventstore.Mapping[user.User](raw)
		if err != nil {
			return e.Wrap(e.EventMappingErr, err)
		}
		*entity = *(&event)
		return nil
	})

	store.SetReplayCase(user.Logout, func(raw any, entity *user.User) error {
		return e.Wrap(e.ResourceDoesNotExist, e.FindResultIsNull)
	})

	store.SetReplayCase(user.NameChanged, func(raw any, entity *user.User) error {
		event, err := eventstore.Mapping[user.UsernameChanged](raw)
		if err != nil {
			return e.Wrap(e.EventMappingErr, err)
		}
		if entity.Name.String() != event.OldName {
			return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
		}
		entity.Name = user.UserName(event.NewName)
		return nil
	})

	store.SetReplayCase(user.PasswordReset, func(raw any, entity *user.User) error {
		event, err := eventstore.Mapping[user.UserPasswordReset](raw)
		if err != nil {
			return e.Wrap(e.EventMappingErr, err)
		}
		if entity.Password.String() != event.OldPassowrd {
			return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
		}
		entity.Password = user.Password(event.NewPassword)
		return nil
	})

	store.SetReplayCase(user.RolesChanged, func(raw any, entity *user.User) error {
		event, err := eventstore.Mapping[user.UserRolesChanged](raw)
		if err != nil {
			return e.Wrap(e.EventMappingErr, err)
		}
		if uint16(entity.Roles) & ^event.OldRights == 0 {
			return e.Wrap(e.EventDisorder, errors.New("data event stream mismatch"))
		}
		entity.Roles = user.UserRole(event.NewRights)
		return nil
	})
}
