package command

import (
	"context"
	"go-blog-ddd/internal/adapter/e"
	"go-blog-ddd/internal/adapter/mdtohtml"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/domain/services"
	"go-blog-ddd/internal/domain/values"
)

type ModifyPostHandler struct {
	repo    aggregates.PostRepository
	service *services.CategoryDomainService
}

func NewModifyPostHandler(repo aggregates.PostRepository, service *services.CategoryDomainService) ModifyPostHandler {
	if repo == nil {
		panic("missing PostRepository")
	}
	if service == nil {
		panic("missing domain service")
	}
	return ModifyPostHandler{repo: repo, service: service}
}

func (s ModifyPostHandler) Handle(ctx context.Context, cmd commands.ModifyPost) error {
	find, err := s.repo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if cmd.Title != nil {
		err := s.setTitle(find, *cmd.Title)
		if err != nil {
			return err
		}
	}

	if cmd.Desc != nil {
		err := s.setDesc(find, []byte(*cmd.Desc))
		if err != nil {
			return err
		}
	}

	if cmd.Visible != nil {
		err := s.setVisible(find, *cmd.Visible)
		if err != nil {
			return err
		}
	}
	if cmd.Tags != nil {
		err := s.setTag(find, *cmd.Tags)
		if err != nil {
			return err
		}
	}
	if cmd.CategoryID != nil {
		err := s.setCategoryID(ctx, find, *cmd.CategoryID)
		if err != nil {
			return err
		}
	}

	return s.repo.Save(ctx, find)
}

func (s ModifyPostHandler) setTag(post *aggregates.Post, tags []string) error {
	tagsLen := len(tags)
	if tagsLen > 4 {
		return e.DErrInvalidOperation.WithMessage("Post最多拥有4个Tag")
	}
	var tagValues = make([]values.Tag, tagsLen)

	for i, tag := range tags {
		newTag, err := values.NewTag(tag)
		if err != nil {
			return e.DErrInvalidOperation.WithMessage(err.Error())
		}
		tagValues[i] = newTag
	}
	post.ModifyPostTags(tagValues)
	return nil
}

func (s ModifyPostHandler) setTitle(post *aggregates.Post, title string) error {
	titleValue, err := values.NewPostTitle(title)
	if err != nil {
		return e.DErrInvalidOperation.WithMessage(err.Error())
	}

	post.ModifyPostTitle(titleValue)
	return nil
}

func (s ModifyPostHandler) setCategoryID(ctx context.Context, post *aggregates.Post, categoryID uint32) error {
	return s.service.ResetCategoryForPost(ctx, post, categoryID)
}

func (s ModifyPostHandler) setDesc(post *aggregates.Post, desc []byte) error {
	descHtml, err := mdtohtml.Convert(desc)
	if err != nil {
		return err
	}
	post.ModifyPostDesc(descHtml)
	return nil
}

func (s ModifyPostHandler) setVisible(post *aggregates.Post, visible bool) error {
	err := post.ModifyVisible(visible)
	if err != nil {
		return e.DErrInvalidOperation.WithMessage(err.Error())
	}
	return nil
}
