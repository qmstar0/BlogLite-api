package util

import (
	"go-blog-ddd/internal/application/query"
	"strings"
	"text/template"
)

var categoryViewTemplate = `
	ID: {{ .ID }}
	Name: {{ .Name }}
	Desc: {{ .Desc }}
	Num: {{ .Num }}
`

var templateObj, _ = template.New("Category View").Parse(categoryViewTemplate)

func CategoryView(view query.CategoryView) (string, error) {
	var s strings.Builder
	err := templateObj.Execute(&s, view)
	if err != nil {
		return "", err
	}
	return s.String(), nil
}
