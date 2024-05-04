package command

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/domain/services"
)

type ResetPostCategoryHandler struct {
	repo    aggregates.PostRepository
	service *services.CategoryDomainService
}

func NewResetPostCategoryHandler(repository aggregates.PostRepository, service *services.CategoryDomainService) ResetPostCategoryHandler {
	if repository == nil {
		panic("missing PostRepository")
	}
	if service == nil {
		panic("missing domain service")
	}
	return ResetPostCategoryHandler{repo: repository, service: service}
}

func (m ResetPostCategoryHandler) Handle(ctx context.Context, cmd commands.ResetPostCategory) error {
	find, err := m.repo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	err = m.service.ResetCategoryForPost(ctx, find, cmd.CategoryID)
	if err != nil {
		return err
	}
	return m.repo.Save(ctx, find)
}
