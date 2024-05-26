package commandhandler

import (
	"context"
	"github.com/qmstar0/domain/internal/domain/aggregates"
	"github.com/qmstar0/domain/internal/domain/commands"
	"github.com/qmstar0/domain/internal/pkg/mdtohtml"
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

	htmlStr, err := mdtohtml.Convert(cmd.MDFile)
	if err != nil {
		return err
	}

	find.ResetContent(htmlStr)
	return m.repo.Save(ctx, find)
}
