package command

import (
	"context"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/adapter/utils/mdtohtml"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"os"
	"path/filepath"
)

type ResetPostContentHandler struct {
	repo aggregates.PostRepository
}

func NewResetPostContentHandler(repo aggregates.PostRepository) ResetPostContentHandler {
	if repo == nil {
		panic("missing PostRepository")
	}
	return ResetPostContentHandler{repo: repo}
}

func (m ResetPostContentHandler) Handle(ctx context.Context, cmd commands.ResetPostContent) error {
	find, err := m.repo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	fileContent, err := os.ReadFile(filepath.Join(config.Conf.Resource.Static.PostFromPath, cmd.MDFilePath))
	if err != nil {
		return err
	}

	htmlStr, err := mdtohtml.Convert(fileContent)
	if err != nil {
		return err
	}

	find.ResetContent(cmd.MDFilePath, htmlStr)

	return m.repo.Save(ctx, find)
}
