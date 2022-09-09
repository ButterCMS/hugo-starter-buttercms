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
			Tag:    tag,
			Layout: "tag",
		}, filepath.Join(contentFolderPath, fmt.Sprintf("%s.md", tag.Slug)))
	}

	CreateFile(data, dataFilePath)
}
