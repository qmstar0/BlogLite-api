package mdtohtml

import (
	"bytes"
	"errors"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
	"strings"
)

type Metadata map[string]string

type ParsedDoc struct {
	Metadata Metadata
	Content  string
}

var defaultMarkdownToHTMLTool = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("github"),
		),
		extension.CJK,
	),
	goldmark.WithParserOptions(
		parser.WithAttribute(),
	),

	goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithHardWraps(),
	),
)

func Convert(mdcontent string) (string, error) {
	var buf bytes.Buffer
	err := defaultMarkdownToHTMLTool.Convert([]byte(mdcontent), &buf)
	if err != nil {
		return "", errors.New("markdown转html失败")
	}
	return buf.String(), nil
}

func Parse(mdcontent string) (*ParsedDoc, error) {
	var err error
	m := make(map[string]string)
	parts := strings.SplitN(mdcontent, "---", 3)

	if len(parts) == 3 {
		err = yaml.Unmarshal([]byte(parts[1]), m)
		if err != nil {
			return nil, err
		}
		return &ParsedDoc{
			Metadata: m,
			Content:  parts[2],
		}, nil
	} else {
		return &ParsedDoc{
			Metadata: m,
			Content:  parts[0],
		}, nil
	}
}
