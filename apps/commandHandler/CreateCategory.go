package commandHandler

import (
	"blog/domain/aggregate/categorys"
	"blog/domain/aggregate/usersRole"
	"blog/domain/common"
	"encoding/json"
	"errors"
	"github.com/ThreeDotsLabs/watermill/message"
)

type CreateCategoryCommandHandler struct {
	categoryRepo categorys.CategoryRepository
	userRoleRepo usersRole.UserRoleRepository
}

var (
	PermissionDeniedErr = errors.New("没有权限")
)

func (h *CreateCategoryCommandHandler) Handler(publisher common.DomainEventPublisher) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		var (
			err error
			cmd categorys.CreateCategoryCommand
		)

		if err = json.Unmarshal(msg.Payload, &cmd); err != nil {
			return err
		}

		userRole, err := h.userRoleRepo.FindByUid(cmd.Uid)
		if !userRole.IsAdmin() {
			return PermissionDeniedErr
		}

		category, err := categorys.CreateCategory(publisher, cmd)
		if err != nil {
			return err
		}
		return h.categoryRepo.Save(category)
	}
}
