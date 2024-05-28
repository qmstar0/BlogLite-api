package commandhandler

import (
	"context"
	"github.com/qmstar0/nightsky-api/internal/domain/aggregates"
	"github.com/qmstar0/nightsky-api/internal/domain/commands"
	"github.com/qmstar0/nightsky-api/internal/domain/services"
	"github.com/qmstar0/nightsky-api/internal/domain/values"
	"github.com/qmstar0/nightsky-api/internal/pkg/e"
	"github.com/qmstar0/nightsky-api/internal/pkg/mdtohtml"
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
	var tagValues = make([]values.Tag, len(tags))

	for i, tag := range tags {
		newTag, err := values.NewTag(tag)
		if err != nil {
			return e.DErrInvalidOperation.WithMessage(err.Error())
		}
		tagValues[i] = newTag
	}
	if err := post.ModifyPostTags(tagValues); err != nil {
		return err
	}
	return nil
}

func (s ModifyPostHandler) setTitle(post *aggregates.Post, title string) error {
	titleValue, err := values.NewPostTitle(title)
	if err != nil {
		return err
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
	if err := post.ModifyVisible(visible); err != nil {
		return err
	}
	return nil
}
