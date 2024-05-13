package mdtohtml

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var defaultMarkdownToHTMLTool = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
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
		return "", err
	}
	return buf.String(), nil
}
