package command

import (
	"context"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/adapter/utils/mdtohtml"
	"go-blog-ddd/internal/application/e"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/domain/values"
	"os"
	"path/filepath"
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

	if find, err := c.postRepo.FindByUri(ctx, uri); err != nil {
		return 0, err
	} else if find != nil {
		return 0, e.ResourceAlreadyExists
	}

	fileContent, err := os.ReadFile(filepath.Join(config.Conf.Resource.Static.PostFromPath, cmd.MDFilePath))
	if err != nil {
		return 0, err
	}

	htmlStr, err := mdtohtml.Convert(fileContent)
	if err != nil {
		return 0, err
	}

	nextID := c.postRepo.NextID(ctx)
	newPost := aggregates.NewPost(nextID, uri, cmd.MDFilePath, htmlStr)
	if err = c.postRepo.Save(ctx, newPost); err != nil {
		return 0, err
	}
	return nextID, nil
}
