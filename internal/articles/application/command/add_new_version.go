package command

import (
	"context"
	"fmt"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/common/constant"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type AddNewVersion struct {
	Uri    string
	Source string
}

type AddNewVersionHandler struct {
	repo        articles.ArticleRepository
	mdParseSer  MarkdownParseService
	dupCheckSer ArticleVersionDuplicationCheckService
}

func NewAddNewVersionHandler(repo articles.ArticleRepository, mdParseSer MarkdownParseService, dupCheckSer ArticleVersionDuplicationCheckService) *AddNewVersionHandler {
	return &AddNewVersionHandler{repo: repo, mdParseSer: mdParseSer, dupCheckSer: dupCheckSer}
}

func (h AddNewVersionHandler) Handle(ctx context.Context, cmd AddNewVersion) error {
	if err := h.checkFileSize(cmd.Source); err != nil {
		return err
	}
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}

	return h.repo.UpdateArticle(
		ctx,
		uri,
		func(article *articles.Article) (*articles.Article, error) {
			version, err := h.mdParseSer.ParseToArticleVersion(cmd.Source)
			if err != nil {
				return nil, err
			}

			if err = h.dupCheckSer.CheckDuplication(ctx, version.Hash); err != nil {
				return nil, err
			}

			if err = article.AddNewVresion(version); err != nil {
				return nil, err
			}
			return article, nil
		},
	)
}

func (h AddNewVersionHandler) checkFileSize(content string) error {
	const maxSize = 1024 * 1024 * constant.MarkdownMaxSizeMB

	if len([]byte(content)) > maxSize {
		return e.InvalidActionError(fmt.Sprintf("上传文件不应超过%dMB", constant.MarkdownMaxSizeMB))
	}
	return nil
}
