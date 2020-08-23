package main

import (
	"os"
	"strings"

	"howett.net/plist"
)

var BookmarkFilePath string

func Parse(path string) Root {
	file, error := os.Open(path)
	if error != nil {
		panic(error)
	}
	decoder := plist.NewDecoder(file)
	root := Root{}
	decoder.Decode(&root)
	return root
}

func Transform(root Root, keyword string) []BookmarkItem {
	lowerKeyword := strings.ToLower(keyword)
	result := []BookmarkItem{}
	for _, folder := range root.Children {
		for _, bookmark := range folder.Children {
			title := bookmark.URIDictionary["title"]
			url := bookmark.URLString
			lowerTitle := strings.ToLower(title)
			lowerURL := strings.ToLower(url)
			if keyword == "" || strings.Contains(lowerTitle, lowerKeyword) || strings.Contains(lowerURL, lowerKeyword) {
				result = append(result, BookmarkItem{title, url})
			}
		}
	}
	return result
}
