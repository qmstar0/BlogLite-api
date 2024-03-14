package command

import (
	"categorys/domain/v1/post"
	"common/e"
	"common/handler"
	"context"
)

type ModifyPostTags struct {
	ID   uint32
	Tags []string
}

type ModifyPostTagsHandler handler.CommandHandler[ModifyPostTags]

type modifyPostTagsHandler struct {
	postRepo post.PostRepository
}

func NewModifyPostTagsHandler(repository post.PostRepository) ModifyPostTagsHandler {
	return &modifyPostTagsHandler{postRepo: repository}
}

func (m modifyPostTagsHandler) Handle(ctx context.Context, cmd ModifyPostTags) error {
	group, err := post.NewTagGroup(cmd.Tags)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	find, err := m.postRepo.Find(ctx, cmd.ID)
	if err != nil {
		return e.Wrap(e.FindEntityErr, err)
	}

	find.SetTags(group)

	return m.postRepo.Save(ctx, find)
}
