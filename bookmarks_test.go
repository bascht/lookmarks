package main

import (
	"testing"
	"log"
)

var bookmarks map[string]string

func init() {
	log.Println("Init")
	bookmarks = buildBookmarks(readBookmarks("test/bookmarks.yaml").Bookmarks, "", "")
	log.Println(bookmarks)
}

func TestBookmarkIndexes(t *testing.T) {
	if len(bookmarks) != 15 {
		t.Error("Expected it to have some bookmarks")
	}
}

func TestBookmarkContent(t *testing.T) {
	if bookmarks["bascht"] != "https://bascht.com" {
		t.Error("Simple URLs are built")
	}
	if bookmarks["bascht-test2-testnested"] != "https://bascht.com/test2/testnested" {
		t.Error("Nested URLs are built")
	}
}

func TestYamlBuilding(t *testing.T) {
	if bookmarks["gitlab-issues-mine"] != "https://gitlab.com/gitlab-org/gitlab/issues/?scope=all&state=opened&assignee_username=bascht" {
		t.Logf(bookmarks["gitlab-issues-mine"])
		t.Error("Simple URLs are built")
	}
}

func TestEmptyKeys(t *testing.T) {
	if bookmarks["bascht-testwithout-anykeys"] != "https://bascht.com/testwithout/anykeys" {
		t.Logf(bookmarks["bascht-testwithout-anykeys"])
		t.Error("Simple URLs are built")
	}
}
