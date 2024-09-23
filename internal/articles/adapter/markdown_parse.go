package adapter

import (
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/pkg/mdtohtml"
	"github.com/qmstar0/BlogLite-api/pkg/utils"
	"strings"
)

type MarkdownParser struct {
}

func NewMarkdownParser() *MarkdownParser {
	return &MarkdownParser{}
}

func (m MarkdownParser) ParseToArticleVersion(content string) (articles.Version, error) {
	parse, err := mdtohtml.Parse(content)
	if err != nil {
		return articles.Version{}, e.InternalServiceError(err.Error())
	}

	text, err := mdtohtml.Convert(parse.Content)
	if err != nil {
		return articles.Version{}, e.InternalServiceError(err.Error())
	}

	description, err := mdtohtml.Convert(strings.TrimSpace(parse.Metadata["description"]))
	if err != nil {
		return articles.Version{}, e.InternalServiceError(err.Error())
	}

	return articles.NewVersion(
		parse.Metadata["title"],
		description,
		text,
		utils.ShortHash(content),
		content,
		parse.Metadata["note"],
	)
}
