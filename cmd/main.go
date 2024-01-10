package main

import (
	"blog/apps/commandHandler"
	"blog/cmd/initCmd"
	"blog/infrastructure/repositoryImpl"
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"time"
)

var duration = time.Second * 10

func main() {

	dao := initCmd.NewDao()
	categoryRepoImpl := repositoryImpl.CategoryRepositoryImpl{DB: dao}
	//userRepoImpl := repositoryImpl.UserRepositoryImpl{DB: dao}
	userRoleRepoImpl := repositoryImpl.UserRoleRepositoryImpl{DB: dao}

	pubsub := gochannel.NewGoChannel(gochannel.Config{}, nil)
	marshaler := cqrs.JSONMarshaler{
		NewUUID:      watermill.NewUUID,
		GenerateName: nil,
	}

	backend := initCmd.NewPubSubBackend(pubsub)

	commandBus := initCmd.NewCommandBus(pubsub, marshaler)
	eventBus := initCmd.NewEventBus(pubsub, marshaler)
	router := initCmd.NewCommandRouter()

	commandProcessor := initCmd.NewCommandProcessor(pubsub, router, marshaler)

	domainEventPublisher := initCmd.NewDomainEventPublisher(eventBus)
	handlers := commandHandler.CategoryCommandHandler{
		Publisher: domainEventPublisher,
		Backend:   backend,
	}

	if err := commandProcessor.AddHandlers(
		handlers.Create(userRoleRepoImpl, categoryRepoImpl),
		handlers.Delete(userRoleRepoImpl, categoryRepoImpl),
		handlers.Update(userRoleRepoImpl, categoryRepoImpl),
	); err != nil {
		panic(err)
	}

	//初始化http服务
	adapter := initCmd.NewHttpAdapter(commandBus, backend)

	go func() {
		//当命令路由启动后， 启动http服务
		<-router.Running()
		err := adapter.StartRun()
		if err != nil {
			panic(err)
		}
	}()

	//启动
	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}
