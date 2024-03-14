package command

import (
	"categorys/domain/v1/post"
	"common/e"
	"common/handler"
	"context"
)

type PublishPost struct {
	ID    uint32
	State bool
}

type PublishPostHandler handler.CommandHandler[PublishPost]

type publishPostHandler struct {
	postRepo post.PostRepository
}

func NewPublishPostHandler(repository post.PostRepository) PublishPostHandler {
	return &publishPostHandler{postRepo: repository}
}

func (p publishPostHandler) Handle(ctx context.Context, cmd PublishPost) error {
	find, err := p.postRepo.Find(ctx, cmd.ID)
	if err != nil {
		return e.Wrap(e.FindEntityErr, err)
	}

	if cmd.State {
		find.Publish()
	} else {
		find.Trash()
	}

	return p.postRepo.Save(ctx, find)
}
