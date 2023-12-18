package articles

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"strings"
)

type ServiceArticle struct {
}

func NewServiceArticle() *ServiceArticle {
	return &ServiceArticle{}
}

// MarkdownToHTML markdown转html
func MarkdownToHTML(markdown string) (string, error) {
	policy := bluemonday.UGCPolicy()
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // 启用 GitHub Flavored Markdown 扩展
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // 自动生成标题的 ID
		),
		goldmark.WithRendererOptions(
			//html.WithHardWraps(), // 处理硬换行
			html.WithUnsafe(),
		),
	)

	var buf strings.Builder
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}

	// 清除潜在的危险内容
	sanitizedHTML := buf.String()
	sanitizedHTML = policy.Sanitize(sanitizedHTML)

	return sanitizedHTML, nil
}
