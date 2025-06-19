package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/akmanon/go-practice-bytest/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post1
Description: description 1
Tags: html, css, js
---
This is a 
Post`
		SecondBody = `Title: Post2
Description: description 2
Tags: java, golang
---
T
i
t
P`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-wolrd2.md": {Data: []byte(SecondBody)},
	}
	posts, err := blogposts.NewBlogPosts(fs)
	if err != nil {
		t.Fatal("Error not Expected")
	}
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post1",
		Description: "description 1",
		Tags:        []string{"html", "css", "js"},
		Body: `This is a 
Post`,
	})

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
