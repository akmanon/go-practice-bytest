package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSep       = "Title: "
	descriptionSep = "Description: "
	tagsSep        = "Tags: "
)

func getPosts(fileSystem fs.FS, f string) (Post, error) {
	postFile, err := fileSystem.Open(f)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return NewPost(postFile)

}

func NewPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	fieldScan := func(sep string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), sep)
	}

	return Post{
		Title:       fieldScan(titleSep),
		Description: fieldScan(descriptionSep),
		Tags:        strings.Split(fieldScan(tagsSep), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
