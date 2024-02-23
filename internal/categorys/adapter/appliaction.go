package adapter

import (
	"blog/pkg/cqrs"
	"blog/pkg/postgresql"
	"categorys/application"
	"categorys/application/command"
	"categorys/application/event"
)

func NewApp(bus *cqrs.Bus) *application.App {
	db := postgresql.GetDB()
	categoryStore := NewCategoryEventStore(db)

	app := &application.App{
		Commands: application.Commands{
			CreateCategory: command.NewCreateCategoryHandler(NewCategoryRepository(), bus),
		},

		Queries: application.Queries{},

		Events: application.Events{
			CategoryCreated: event.NewCategoryCreatedHandler(categoryStore),
		},
	}

	addHandlerToBus(bus, app)
	return app
}

func addHandlerToBus(bus *cqrs.Bus, app *application.App) {
	events := app.Events

	bus.AddHandler(
		cqrs.NewHandler[event.CategoryCreated](events.CategoryCreated.Handle),
	)
}
