package main

import (
	"log"
	"os"

	"github.com/akmanon/go-practice-bytest/blogposts"
)

func main() {
	posts, err := blogposts.NewBlogPosts(os.DirFS(""))
	if err != nil {
		log.Fatal(err)

	}
	log.Println(posts)
}
