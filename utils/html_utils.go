package utils

import (
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHtml(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func ExtractH1(html []byte) (string, string) {
	h1 := regexp.MustCompile(`^<h1>.*</h1>`)
	title := strings.Split(string(h1.Find(html)), ">")[1]
	title = strings.Split(title, "<")[0]
	content := string(h1.ReplaceAll(html, nil))
	return title, content
}
