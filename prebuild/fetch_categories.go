package main

import ButterCMS "github.com/ButterCMS/buttercms-go"

type Category struct {
	Name string
	Slug string
}

func FetchCategories(path string) {
	categoriesResponse, err := ButterCMS.GetCategories(map[string]string{})
	HandleErr(err)

	data := []Category{}
	for _, category := range categoriesResponse.CategoryList {
		data = append(data, Category{
			Name: category.Name,
			Slug: category.Slug,
		})
	}

	CreateFile(data, path)
}
