package command

import (
	"context"
	"errors"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/domain/values"
)

type ModifyPostTagsHandler struct {
	repo aggregates.PostRepository
}

func NewModifyPostTagsHandler(
	repo aggregates.PostRepository,
) ModifyPostTagsHandler {
	if repo == nil {
		panic("missing PostRepository")
	}
	return ModifyPostTagsHandler{repo: repo}
}

func (m ModifyPostTagsHandler) Handle(ctx context.Context, cmd commands.ModifyPostTags) error {
	tagsLen := len(cmd.NewTags)
	if tagsLen > 4 {
		return errors.New("Post最多拥有4个Tag")
	}

	var tags = make([]values.Tag, tagsLen)
	for i, tag := range cmd.NewTags {
		newTag, err := values.NewTag(tag)
		if err != nil {
			return err
		}
		tags[i] = newTag
	}

	find, err := m.repo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	find.ModifyPostTags(tags)

	return m.repo.Save(ctx, find)
}
