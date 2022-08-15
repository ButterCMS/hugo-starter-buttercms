package main

import (
	"fmt"
	"path/filepath"

	ButterCMS "github.com/ButterCMS/buttercms-go"
)

func FetchBlogPosts(pathToFiles string) {
	response, err := ButterCMS.GetPosts(map[string]string{})
	HandleErr(err)

	for _, post := range response.PostList {
		CreateFile(post, filepath.Join(pathToFiles, fmt.Sprintf("%s.md", post.Slug)))
	}
}
