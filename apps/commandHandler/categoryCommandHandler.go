package commandHandler

import (
	"blog/apps/commandResult"
	"blog/domain/aggregate/categorys"
	"blog/domain/aggregate/userRole"
	"blog/domain/common"
	"context"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

type CategoryCommandHandler struct {
	Publisher common.DomainEventPublisher
	Backend   requestreply.Backend[commandResult.StateCode]
}

func (c CategoryCommandHandler) Create(UserRoleRepo userRole.UserRoleRepository, CategoryRepo categorys.CategoryRepository) cqrs.CommandHandler {
	return requestreply.NewCommandHandlerWithResult[categorys.CreateCategoryCommand](
		"Handler.Command.Category.Create",
		c.Backend,
		func(ctx context.Context, cmd *categorys.CreateCategoryCommand) (commandResult.StateCode, error) {
			role, err := UserRoleRepo.FindByUid(cmd.Uid)
			if err != nil {
				return commandResult.ItemNotExist, nil
			}
			if !role.IsAdmin() {
				return commandResult.PermissionDenied, nil
			}
			category, err := categorys.CreateCategory(c.Publisher, *cmd)
			if err != nil {
				return commandResult.EventPublishErr, err
			}

			err = CategoryRepo.Save(category)
			if err != nil {
				return commandResult.DBSaveErr, err
			}

			return commandResult.Successed, nil
		})
}

func (c CategoryCommandHandler) Update(UserRoleRepo userRole.UserRoleRepository, CategoryRepo categorys.CategoryRepository) cqrs.CommandHandler {
	return requestreply.NewCommandHandlerWithResult[categorys.UpdateCategoryCommand](
		"Handler.Command.Category.Update",
		c.Backend,
		func(ctx context.Context, cmd *categorys.UpdateCategoryCommand) (commandResult.StateCode, error) {
			role, err := UserRoleRepo.FindByUid(cmd.Uid)
			if err != nil {
				return commandResult.ItemNotExist, nil
			}

			if !role.IsAdmin() {
				return commandResult.PermissionDenied, nil
			}

			category, err := CategoryRepo.FindById(cmd.CategoryId)

			err = category.Update(c.Publisher, *cmd)
			if err != nil {
				return commandResult.EventPublishErr, nil
			}

			err = CategoryRepo.Save(category)
			if err != nil {
				return commandResult.DBSaveErr, err
			}

			return commandResult.Successed, nil
		},
	)
}

func (c CategoryCommandHandler) Delete(UserRoleRepo userRole.UserRoleRepository, CategoryRepo categorys.CategoryRepository) cqrs.CommandHandler {
	return requestreply.NewCommandHandlerWithResult[categorys.DeleteCategoryCommand](
		"Handler.Command.Category.Delete",
		c.Backend,
		func(ctx context.Context, cmd *categorys.DeleteCategoryCommand) (commandResult.StateCode, error) {
			role, err := UserRoleRepo.FindByUid(cmd.Uid)
			if err != nil {
				return commandResult.ItemNotExist, nil
			}

			if !role.IsAdmin() {
				return commandResult.PermissionDenied, nil
			}

			category, err := CategoryRepo.FindById(cmd.CategoryId)

			err = category.Delete(c.Publisher, *cmd)
			if err != nil {
				return commandResult.EventPublishErr, nil
			}

			err = CategoryRepo.Save(category)
			if err != nil {
				return commandResult.DBSaveErr, err
			}

			return commandResult.Successed, nil
		},
	)
}
