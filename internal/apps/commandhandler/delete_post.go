package commandhandler

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
)

type DeletePostHandler struct {
	postRepo aggregates.PostRepository
}

func NewDeletePostHandler(repository aggregates.PostRepository) DeletePostHandler {
	if repository == nil {
		panic("missing PostRepository")
	}
	return DeletePostHandler{postRepo: repository}
}

func (d DeletePostHandler) Handle(ctx context.Context, cmd commands.DeletePost) error {
	return d.postRepo.Delete(ctx, cmd.ID)
}
