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

const dataPath = "./remote/data"

const contentPath = "./remote/content"

func fetch() {
	ButterCMS.SetAuthToken(os.Getenv("BUTTERCMS_API_TOKEN"))

	createFolderIfNotExists(dataPath)
	createFolderIfNotExists(contentPath)

	CreateHeaderFile(filepath.Join(dataPath, "header.json"))

	landingPagesPath := filepath.Join(contentPath, "landing_page")
	createFolderIfNotExists(landingPagesPath)
	CreateLandingPagesFiles(landingPagesPath)
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
		fmt.Println("Error fetching data: %w", err)

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
