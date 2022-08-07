package main

import (
	ButterCMS "github.com/ButterCMS/buttercms-go"
)

type NavItem struct {
	Url   string
	Label string
}

type HeaderFile struct {
	DefaultActiveItemUrl string
	NavItems             []NavItem
}

func CreateHeaderFile(path string) {
	data := HeaderFile{
		DefaultActiveItemUrl: "#home",
		NavItems:             getNavigationItems(),
	}

	CreateFile(data, path)
}

func getNavigationItems() []NavItem {
	navItemsResponse, err := ButterCMS.GetContentFields([]string{"navigation_menu_item"}, map[string]string{})
	HandleErr(err)

	var navItems []NavItem
	for _, element := range navItemsResponse.Data["navigation_menu_item"].([]interface{}) {
		casted := element.(map[string]interface{})

		navItem := NavItem{}
		if url, ok := casted["url"].(string); ok {
			navItem.Url = url
		}
		if label, ok := casted["label"].(string); ok {
			navItem.Label = label
		}

		navItems = append(navItems, navItem)
	}

	return navItems
}
