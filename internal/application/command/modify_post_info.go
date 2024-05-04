package command

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/domain/values"
)

type ModifyPostInfoHandler struct {
	repo aggregates.PostRepository
}

func NewModifyPostInfoHandler(repo aggregates.PostRepository) ModifyPostInfoHandler {
	if repo == nil {
		panic("missing PostRepository")
	}
	return ModifyPostInfoHandler{repo: repo}
}

func (s ModifyPostInfoHandler) Handle(ctx context.Context, cmd commands.ModifyPostInfo) error {
	title, err := values.NewPostTitle(cmd.Title)
	if err != nil {
		return err
	}
	find, err := s.repo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	find.ModifyPostInfo(title, cmd.Desc)
	return s.repo.Save(ctx, find)
}
