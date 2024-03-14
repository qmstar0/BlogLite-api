package command

import (
	"categorys/domain/v1/post"
	"common/e"
	"common/handler"
	"context"
)

type CreatePost struct {
	UserID uint32
	Uri    string
}

type CreatePostHandler handler.CommandHandler[CreatePost]

type createPostHandler struct {
	postRepo post.PostRepository
}

func NewCreatePostHandler(repository post.PostRepository) CreatePostHandler {
	return createPostHandler{postRepo: repository}
}

func (c createPostHandler) Handle(ctx context.Context, cmd CreatePost) error {
	uri, err := post.NewUri(cmd.Uri)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	if exist, err := c.postRepo.Exist(ctx, uri.ToID()); err != nil {
		return e.Wrap(e.FindEntityErr, err)
	} else if exist {
		return e.Wrap(e.ResourceCreated, e.ResourceAlreadyExists)
	}

	newPost := post.NewPost(cmd.UserID, uri)

	return c.postRepo.Save(ctx, newPost)
}
