package storage

import (
	"pier/notify"
	"pier/reader/models"
	"pier/storage"
)

func Feeds() []*models.Feed {
	feeds := []*models.Feed{}

	db := storage.DB()
	rows, err := db.Query("SELECT `name`, `url`, `disabled`, `updated` FROM `reader_feeds`")
	if err != nil {
		notify.ErrorAlert("reader", "get feeds", err)
		return feeds
	}
	defer rows.Close()
	for rows.Next() {
		var feed models.Feed
		if err := rows.Scan(&feed.Name, &feed.Url, &feed.Disabled, &feed.Updated); err != nil {
			notify.ErrorAlert("reader", "get feeds", err)
		}
		feeds = append(feeds, &feed)
	}

	return feeds
}
