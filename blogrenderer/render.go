package blogrenderer

import (
	"embed"
	"io"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type PostRenderer struct {
	templ        *template.Template
	mdParser     *parser.Parser
	htmlRenderer *html.Renderer
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return &PostRenderer{templ: templ, mdParser: p, htmlRenderer: renderer}, nil
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	doc := r.mdParser.Parse([]byte(post.Body))
	post = Post{
		Title:       post.Title,
		Description: post.Description,
		Body:        string(markdown.Render(doc, r.htmlRenderer)),
		Tags:        post.Tags,
	}
	return r.templ.ExecuteTemplate(w, "blog.gohtml", post)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
