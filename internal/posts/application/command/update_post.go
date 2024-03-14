package command

import (
	"categorys/domain/v1/post"
	"common/e"
	"common/handler"
	"context"
)

type UpdatePost struct {
	ID         uint32
	Title      string
	Content    string
	SeoDesc    string
	Tags       []string
	CategoryID uint32
}

type UpdatePostHandler handler.CommandHandler[UpdatePost]

type updatePostHandler struct {
	postRepo post.PostRepository
}

func NewUpdatePostHandler(repository post.PostRepository) UpdatePostHandler {
	return &updatePostHandler{postRepo: repository}
}

func (u updatePostHandler) Handle(ctx context.Context, cmd UpdatePost) error {
	find, err := u.postRepo.Find(ctx, cmd.ID)
	if err != nil {
		return e.Wrap(e.FindEntityErr, err)
	}

	find.Modify(cmd.Title, cmd.Content, cmd.SeoDesc)

	return u.postRepo.Save(ctx, find)
}
