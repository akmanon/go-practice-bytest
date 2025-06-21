package blogrenderer

import (
	"bytes"
	"io"
	"testing"

	"github.com/akmanon/go-practice-bytest/blogposts"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	apost := blogposts.Post{
		Title:       "hello world",
		Description: "This is a description",
		Tags:        []string{"html", "css", "js"},
		Body: `This
is a
**body**`,
	}
	t.Run("converts a single post into html", func(t *testing.T) {
		buf := bytes.Buffer{}
		postRenderer, err := NewRenderer()
		if err != nil {
			t.Fatal(err)
		}
		err = postRenderer.Render(&buf, apost)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

}

func BenchmarkRender(b *testing.B) {
	apost := blogposts.Post{
		Title:       "hello world",
		Description: "This is a description",
		Tags:        []string{"html", "css", "js"},
		Body: `This
is a
body`,
	}
	postRenderer, _ := NewRenderer()
	for b.Loop() {
		postRenderer.Render(io.Discard, apost)
	}
}
