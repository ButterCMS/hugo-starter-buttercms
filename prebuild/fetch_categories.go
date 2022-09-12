package main

import (
	"fmt"
	"path/filepath"

	ButterCMS "github.com/ButterCMS/buttercms-go"
)

type Category struct {
	Name string
	Slug string
}

type CategoryContentFile struct {
	Category

	Title       string
	Description string

	Layout string `json:"layout"`
}

func FetchCategories(dataFilepath string, contentFolderPath string) {
	categoriesResponse, err := ButterCMS.GetCategories(map[string]string{})
	HandleErr(err)

	data := []Category{}
	for _, category := range categoriesResponse.CategoryList {
		category := Category{
			Name: category.Name,
			Slug: category.Slug,
		}
		data = append(data, category)

		CreateFile(CategoryContentFile{
			Category:    category,
			Title:       fmt.Sprintf("Sample Blog - category: %s", category.Name),
			Description: fmt.Sprintf("Sample blog powered by ButterCMS, showing category: %s", category.Name),

			Layout: "category",
		}, filepath.Join(contentFolderPath, fmt.Sprintf("%s.md", category.Slug)))
	}

	CreateFile(data, dataFilepath)
}
