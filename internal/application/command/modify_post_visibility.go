package command

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
)

type ModifyPostVisibleHandler struct {
	postRepo aggregates.PostRepository
}

func NewModifyPostVisibleHandler(repository aggregates.PostRepository) ModifyPostVisibleHandler {
	if repository == nil {
		panic("missing PostRepository")
	}
	return ModifyPostVisibleHandler{postRepo: repository}
}

func (m ModifyPostVisibleHandler) Handle(ctx context.Context, cmd commands.ModifyPostVisibility) (err error) {
	findPost, err := m.postRepo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if err = findPost.ModifyVisible(cmd.Visible); err != nil {
		return err
	}
	return m.postRepo.Save(ctx, findPost)
}
