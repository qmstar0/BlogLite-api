package mdtohtml

import (
	"bytes"
	"github.com/qmstar0/nightsky-api/internal/pkg/e"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var defaultMarkdownToHTMLTool = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.NewFootnote(
			extension.WithFootnoteIDPrefix("footnote-"),
		),
		highlighting.NewHighlighting(
			highlighting.WithStyle("github"),
		),
		extension.CJK,
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
		parser.WithAttribute(),
	),

	goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithHardWraps(),
	),
)

func Convert(mdstr []byte) (string, error) {
	var buf bytes.Buffer
	err := defaultMarkdownToHTMLTool.Convert(mdstr, &buf)
	if err != nil {
		return "", e.UErrUtilMarkdownToHTML.WithError(err)
	}
	return buf.String(), nil
}
