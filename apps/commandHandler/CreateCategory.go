package commandHandler

import (
	"blog/domain/aggregate/categorys"
	"context"
)

type CreateCategoryCommandHandler struct {
}

func (c CreateCategoryCommandHandler) HandlerName() string {
	return "Handler.Command.Category.Create"
}

func (c CreateCategoryCommandHandler) NewCommand() any {
	return &categorys.CreateCategoryCommand{}
}

func (c CreateCategoryCommandHandler) Handle(ctx context.Context, cmd any) error {
	//TODO implement me
	panic("implement me")
}
