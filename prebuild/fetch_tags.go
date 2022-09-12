package main

import (
	"fmt"
	"path/filepath"

	ButterCMS "github.com/ButterCMS/buttercms-go"
)

type Tag struct {
	Name string
	Slug string
}

type TagContentFile struct {
	Tag

	Title       string
	Description string

	Layout string `json:"layout"`
}

func FetchTags(dataFilePath string, contentFolderPath string) {
	tagsResponse, err := ButterCMS.GetTags(map[string]string{})
	HandleErr(err)

	data := []Tag{}
	for _, tag := range tagsResponse.TagList {
		tag := Tag{
			Name: tag.Name,
			Slug: tag.Slug,
		}
		data = append(data, tag)

		CreateFile(TagContentFile{
			Tag: tag,

			Title:       fmt.Sprintf("Sample Blog - tag: %s", tag.Name),
			Description: fmt.Sprintf("Sample blog powered by ButterCMS, showing tag: %s", tag.Name),

			Layout: "tag",
		}, filepath.Join(contentFolderPath, fmt.Sprintf("%s.md", tag.Slug)))
	}

	CreateFile(data, dataFilePath)
}
