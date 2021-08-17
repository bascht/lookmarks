package main

import (
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Bookmark struct {
	Key string `json:"key,omitempty"`
	Url string `json:url`
	Sub []Bookmark `json:sub`
}

type BookmarkList struct {
	Bookmarks []Bookmark `yaml:"bookmarks"`
}

func buildBookmarks(bookmarks []Bookmark, prefix string, url string) map[string]string {
	var bookmarkMap = make(map[string]string)

	var bookmarkKey string
	var bookmarkUrl string

	for _, bookmark := range bookmarks {
		bookmarkUrl = url + bookmark.Url

		if bookmark.Key != "" {
			bookmarkKey = prefix + bookmark.Key
		} else {
			bookmarkKey = prefix + bookmark.Url
		}

		bookmarkMap[bookmarkKey] = bookmarkUrl

		if len(bookmark.Sub) > 0 {
			mergo.Merge(&bookmarkMap, buildBookmarks(bookmark.Sub, bookmarkKey+"-", bookmarkUrl+"/"))
		}
	}
	return bookmarkMap
}

func readBookmarks(file string) BookmarkList {
	var bookmarkList BookmarkList

	yamlFile, err := ioutil.ReadFile(file)

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &bookmarkList)
	if err != nil {
		log.Fatalf("Failed to read the Yaml file: %v", err)
	}

	return bookmarkList
}
