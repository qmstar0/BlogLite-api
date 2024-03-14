package command

import (
	"categorys/domain/v1/post"
	"common/e"
	"common/handler"
	"context"
)

type ResetPostCategroy struct {
	ID         uint32
	CategoryID uint32
}

type ModifyPostCateogryHandler handler.CommandHandler[ResetPostCategroy]

type modifyPostCategoryHandler struct {
	postRepo post.PostRepository
}

func NewModifyPostCategoryHandler(repository post.PostRepository) ModifyPostCateogryHandler {
	return &modifyPostCategoryHandler{postRepo: repository}
}

func (r modifyPostCategoryHandler) Handle(ctx context.Context, cmd ResetPostCategroy) error {
	find, err := r.postRepo.Find(ctx, cmd.ID)
	if err != nil {
		return e.Wrap(e.FindEntityErr, err)
	}

	find.SetCategory(cmd.CategoryID)
	return r.postRepo.Save(ctx, find)
}
