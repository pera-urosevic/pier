package net

import (
	"pier/api/reader/database/model"

	"github.com/mmcdole/gofeed"
)

func Fetch(feed *model.Feed) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = "Pier 2.0"
	res, err := fp.ParseURL(feed.URL)
	return res, err
}
