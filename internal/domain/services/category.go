package services

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
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
	cate, err := p.cateRepo.FindByIDOrErr(ctx, newCategoryID)
	if err != nil {
		return err
	}
	post.CategoryID = cate.ID
	return nil
}
