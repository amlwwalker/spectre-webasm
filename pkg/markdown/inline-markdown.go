package markdown

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func InlineMarkdown(md string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	res := string(markdown.ToHTML([]byte(md), parser, nil))
	return res
}
