package main

import ButterCMS "github.com/ButterCMS/buttercms-go"

type Tag struct {
	Name string
	Slug string
}

func FetchTags(path string) {
	tagsResponse, err := ButterCMS.GetTags(map[string]string{})
	HandleErr(err)

	data := []Tag{}
	for _, tag := range tagsResponse.TagList {
		data = append(data, Tag{
			Name: tag.Name,
			Slug: tag.Slug,
		})
	}

	CreateFile(data, path)
}
