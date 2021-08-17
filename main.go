package main

import (
	"fmt"
	"github.com/adrg/xdg"
	"github.com/skratchdot/open-golang/open"
	"log"
	"os"
)

func main() {
	bookmarkFile, err := xdg.DataFile("bookmarks.yaml")
	if err != nil {
		log.Printf("Could not retrieve bookmarks: %s", err)
		os.Exit(1)
	}

	var bookmarkList = readBookmarks(bookmarkFile)
	var bookmarks = buildBookmarks(bookmarkList.Bookmarks, "", "")
	var is_preview = os.Getenv("ROFI_PREVIEW") == "1"

	if len(os.Args) == 2 {
		bookmark := os.Args[1]

		if url, ok := bookmarks[bookmark]; ok {
			if is_preview {
				fmt.Println(url)
			} else {
				fmt.Printf("Going to open %s", url)
				err := open.Run(url)
				if err != nil {
					os.Exit(1)
				}
			}
		} else {
			log.Printf("Could not find %s in bookmark list", bookmark)
			os.Exit(1)
		}
		os.Exit(0)
	}

	for key := range bookmarks {
		fmt.Println(key)
	}
}
