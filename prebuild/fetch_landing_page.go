package main

import (
	"fmt"
	"path/filepath"

	ButterCMS "github.com/ButterCMS/buttercms-go"
)

const defaultLandingPageSlug = "landing-page-with-components"

type LandingPageSectionBase struct {
	ScrollAnchorId string
	Headline       string
	SubHeadline    string

	IsSet bool
}

type HeroSection struct {
	LandingPageSectionBase

	ButtonUrl   string
	ButtonLabel string
	Image       string
}

type SEOMetadata struct {
	Title       string
	Description string
}

type LandingPageFile struct {
	SEOMetadata

	HeroSection HeroSection
}

func CreateLandingPagesFiles(pathToFiles string) {
	pageType := "landing-page"

	response, pageFetchingErr := ButterCMS.GetPages(pageType, map[string]string{})
	HandleErr(pageFetchingErr)

	for _, page := range response.PageList {
		createLandingPageFile(pathToFiles, page)
	}
}

func createLandingPageFile(pathToFile string, page ButterCMS.Page) {
	data := LandingPageFile{
		SEOMetadata: parseSEOMetadata(page),
	}

	if body, err := GetValue[[]interface{}](page.Fields, "body"); err == nil {
		for _, untypedSection := range body {
			section := untypedSection.(map[string]interface{})

			scrollAnchorId, _ := getSectionFieldsValue[string](section, "scroll_anchor_id")
			headline, _ := getSectionFieldsValue[string](section, "headline")

			switch scrollAnchorId {
			case "#home":
				buttonLabel, _ := getSectionFieldsValue[string](section, "button_label")
				buttonUrl, _ := getSectionFieldsValue[string](section, "button_url")
				image, _ := getSectionFieldsValue[string](section, "image")
				subHeadline, _ := getSectionFieldsValue[string](section, "subheadline")

				data.HeroSection.Headline = headline
				data.HeroSection.ScrollAnchorId = scrollAnchorId
				data.HeroSection.ButtonLabel = buttonLabel
				data.HeroSection.ButtonUrl = buttonUrl
				data.HeroSection.Image = image
				data.HeroSection.SubHeadline = subHeadline

				data.HeroSection.IsSet = true
			}

		}
	}

	// fmt.Printf("\n page: %v", data)

	if defaultLandingPageSlug == page.Slug {
		CreateFile(data, filepath.Join(pathToFile, "_index.md"))
	}

	CreateFile(data, filepath.Join(pathToFile, fmt.Sprintf("%s.md", page.Slug)))
}

func parseSEOMetadata(page ButterCMS.Page) SEOMetadata {
	result := SEOMetadata{}

	if seo, err := GetValue[map[string]interface{}](page.Fields, "seo"); err == nil {
		title, _ := GetValue[string](seo, "title")
		description, _ := GetValue[string](seo, "description")

		result.Title = title
		result.Description = description
	}

	return result
}

func getSectionFieldsValue[T any](input map[string]interface{}, name string) (T, error) {
	fields, err := GetValue[map[string]interface{}](input, "fields")
	if err != nil {
		var null T
		return null, err
	}

	return GetValue[T](fields, name)
}
