package adapter

import (
	"blog/internal/categorys/application"
	"blog/internal/categorys/application/command"
	"context"
	"fmt"
	"github.com/qmstar0/eio"
	"github.com/qmstar0/eio-cqrs/cqrs"
)

func NewApp(pub eio.Publisher, sub eio.Subscriber, routerBus cqrs.RouterBus) *application.App {
	var err error
	err = routerBus.AddHandlers(
		// Command.CreateCategory
		cqrs.NewHandler[command.CreateCategory](
			"Command.CreateCategory", sub,
			command.NewCreateCategoryHandler(NewCategoryRepository()).Handle),
	)
	if err != nil {
		panic(fmt.Errorf("an error occurred while `NewApp`: %w", err))
	}

	err = routerBus.AddHandlers()
	if err != nil {
		panic(fmt.Errorf("an error occurred while `NewApp`: %w", err))
	}

	CommandPublisher := routerBus.WithPublisher(pub)
	QueryPublisher := routerBus.WithPublisher(pub)
	return &application.App{
		CommandsBus: busAdapter{CommandPublisher},
		QueriesBus:  busAdapter{QueryPublisher},
	}
}

type busAdapter struct {
	bus cqrs.PublishBus
}

func (b busAdapter) Publish(ctx context.Context, v any) error {
	err := b.bus.Publish(ctx, v)
	if err != nil {
		return err
	}
	return nil
}
