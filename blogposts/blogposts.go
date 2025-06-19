package blogposts

import (
	"io/fs"
)

func NewBlogPosts(fileSystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for _, f := range dir {
		post, err := getPosts(fileSystem, f.Name())

		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
