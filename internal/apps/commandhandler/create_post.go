package commandhandler

import (
	"context"
	"github.com/qmstar0/nightsky-api/internal/domain/aggregates"
	"github.com/qmstar0/nightsky-api/internal/domain/commands"
	"github.com/qmstar0/nightsky-api/internal/domain/values"
	"github.com/qmstar0/nightsky-api/internal/pkg/mdtohtml"
)

type CreatePostHandler struct {
	postRepo aggregates.PostRepository
}

func NewCreatePostHandler(repository aggregates.PostRepository) CreatePostHandler {
	if repository == nil {
		panic("missing PostRepository")
	}
	return CreatePostHandler{postRepo: repository}
}

func (c CreatePostHandler) Handle(ctx context.Context, cmd commands.CreatePost) (uint32, error) {
	uri, err := values.NewPostUri(cmd.Uri)
	if err != nil {
		return 0, err
	}

	if err = c.postRepo.ResourceUniquenessCheck(ctx, uri); err != nil {
		return 0, err
	}

	htmlStr, err := mdtohtml.Convert(cmd.MDFile)
	if err != nil {
		return 0, err
	}

	nextID, err := c.postRepo.NextID(ctx)
	if err != nil {
		return 0, err
	}

	newPost := aggregates.NewPost(nextID, uri, htmlStr)

	if err = c.postRepo.Save(ctx, newPost); err != nil {
		return 0, err
	}
	return nextID, nil
}
