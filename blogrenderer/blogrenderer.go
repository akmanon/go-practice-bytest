package blogrenderer

import (
	"bytes"
	"embed"
	"io"
	"text/template"

	"github.com/akmanon/go-practice-bytest/blogposts"
	"github.com/yuin/goldmark"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (p *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	if err := p.postBodyParser(&post); err != nil {
		return err
	}

	if err := p.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}
	return nil
}

func (p *PostRenderer) postBodyParser(post *blogposts.Post) error {
	buf := bytes.Buffer{}
	err := goldmark.Convert([]byte(post.Body), &buf)
	post.Body = buf.String()
	if err != nil {
		return err
	}
	return nil
}
