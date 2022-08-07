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

const fileMode = 0644 // TODO check if value is correct

const dataPath = "./remote/data"

func fetch() {
	ButterCMS.SetAuthToken(os.Getenv("BUTTERCMS_API_TOKEN"))

	createFolderIfNotExists(dataPath)

	CreateHeaderFile(filepath.Join(dataPath, "header.json"))
}

func createFolderIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, fileMode)
		if err != nil {
			HandleErr(err)
		}
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
		fmt.Println("Error fetching data: %w", err) // todo change wording

		os.Exit(1)
	}
}
