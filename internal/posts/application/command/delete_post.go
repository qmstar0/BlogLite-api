package command

import (
	"categorys/domain/v1/post"
	"common/e"
	"common/handler"
	"context"
)

type DeletePost struct {
	ID uint32
}

type DeletePostHandler handler.CommandHandler[DeletePost]

type deletePostHandler struct {
	postRepo post.PostRepository
}

func NewDeletePostHandler(repository post.PostRepository) DeletePostHandler {
	return &deletePostHandler{postRepo: repository}
}

func (d deletePostHandler) Handle(ctx context.Context, cmd DeletePost) error {
	find, err := d.postRepo.Find(ctx, cmd.ID)
	if err != nil {
		return e.Wrap(e.FindEntityErr, err)
	}
	find.Delete()
	return d.postRepo.Save(ctx, find)
}
