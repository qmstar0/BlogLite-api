package services

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/pkg/e"
)

type CategoryDomainService struct {
	cateRepo aggregates.CategoryRepository
}

func NewPostDomainService(cateRepo aggregates.CategoryRepository) *CategoryDomainService {
	return &CategoryDomainService{
		cateRepo: cateRepo,
	}
}

func (p CategoryDomainService) ResetCategoryForPost(ctx context.Context, post *aggregates.Post, newCategoryID uint32) error {
	cate, err := p.cateRepo.FindOrErrByID(ctx, newCategoryID)
	if err != nil {
		return e.DErrInvalidOperation.WithError(err).WithMessage("没有该分类")
	}
	post.CategoryID = cate.ID
	return nil
}
