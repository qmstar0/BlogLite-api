package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
)

type MarkdownParseService interface {
	// ParseToArticleVersion 解析将传入的markdown文件内容
	// 将文件元数据中的description和文件正文转为html内容
	ParseToArticleVersion(content string) (articles.Version, error)
}

type CategoryValidityCheckService interface {
	// CategoryExist 检查分类是否真实存在。文章不能被分配在一个不存在的分类下
	CategoryExist(ctx context.Context, categoryID string) error
}

type ArticleVersionDuplicationCheckService interface {
	// CheckDuplication 检查上传的markdown文件hash是否存在。
	// 引入此服务的原因是：一个markdown被上传到不同的两个文章资源下，由于解析结果（version）的hash相同，
	// 这两个文章资源实际持有同一个版本资源，当其中一个执行删除`共享的版本`时，另一个同样会被删除。
	// 这将会造成数据的不一致
	CheckDuplication(ctx context.Context, versionHash string) error
}
