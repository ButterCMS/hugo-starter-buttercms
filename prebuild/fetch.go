package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	ButterCMS "github.com/ButterCMS/buttercms-go"
)

const fileMode = 0777

func fetch(remoteFolder string) {
	apiKey := os.Getenv("BUTTERCMS_API_TOKEN")
	if apiKey == "" {
		HandleErr(errors.New("BUTTERCMS_API_TOKEN env is not set. Get your free API key on https://buttercms.com/join/"))
	}

	previewModeEnv := os.Getenv("BUTTERCMS_PREVIEW")
	previewMode := true
	if previewModeEnv == "false" {
		previewMode = false
	}
	ButterCMS.SetPreviewMode(previewMode)

	fmt.Printf("Preview mode: %t\n", previewMode)

	ButterCMS.SetAuthToken(apiKey)

	fmt.Println("Fetching of content started.")

	dataPath := filepath.Join(remoteFolder, "data")
	contentPath := filepath.Join(remoteFolder, "content")
	landingPagePath := filepath.Join(contentPath, "landing-page")
	blogPath := filepath.Join(contentPath, "blog")
	categoryPath := filepath.Join(blogPath, "category")
	tagPath := filepath.Join(blogPath, "tag")

	removeFolderRecursively(dataPath)
	removeFolderRecursively(contentPath)

	createFolderIfNotExists(dataPath)
	createFolderIfNotExists(contentPath)
	createFolderIfNotExists(landingPagePath)
	createFolderIfNotExists(blogPath)
	createFolderIfNotExists(categoryPath)
	createFolderIfNotExists(tagPath)

	FetchHeader(filepath.Join(dataPath, "header.json"), "#home")
	FetchCategories(filepath.Join(dataPath, "categories.json"), categoryPath)
	FetchTags(filepath.Join(dataPath, "tags.json"), tagPath)
	FetchLandingPages(contentPath, landingPagePath)

	FetchBlogPosts(blogPath)

	fmt.Println("Fetching of content successfully finished.")
}

func createFolderIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, fileMode)

		HandleErr(err)
	}
}

func CreateFile(data any, path string) {
	file, jsonErr := json.MarshalIndent(data, "", " ")
	HandleErr(jsonErr)

	fileErr := ioutil.WriteFile(path, file, fileMode)
	HandleErr(fileErr)
}

func HandleErr(err error) {
	if err != nil {
		if err.Error() == "Unauthorized" {
			fmt.Println("Your ButterCMS token is not valid. Please check your token and try again.")
		} else {
			fmt.Printf("Error fetching data: %v \n", err)
		}

		os.Exit(1)
	}
}

func GetValue[T any](input map[string]interface{}, name string) (T, error) {
	field, ok := input[name].(T)

	if !ok {
		var null T
		return null, fmt.Errorf("unable to get property %s from map", name)
	}

	return field, nil
}

func removeFolderRecursively(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		HandleErr(err)
	}
}
