package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type CheckAndDeleteCategory struct {
	CategorySlug string
}

type CheckAndDeleteCategoryHandler struct {
	ser           GetCategoryUsedService
	deleteHandler *DeleteCategoryHandler
}

func NewCheckAndDeleteCategoryHandler(ser GetCategoryUsedService, deleteHandler *DeleteCategoryHandler) *CheckAndDeleteCategoryHandler {
	return &CheckAndDeleteCategoryHandler{ser: ser, deleteHandler: deleteHandler}
}

func (h *CheckAndDeleteCategoryHandler) Handle(ctx context.Context, cmd CheckAndDeleteCategory) error {
	used, err := h.ser.IsUsed(ctx, cmd.CategorySlug)
	if err != nil {
		return err
	}
	if used {
		return e.InvalidActionError("无法删除该分类：此分类下仍存在关联的文章，请先删除或移除关联的文章")
	}

	return h.deleteHandler.Handle(ctx, DeleteCategory{CategorySlug: cmd.CategorySlug})
}
