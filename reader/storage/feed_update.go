package storage

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func FeedUpdate(feed *model.Feed) {
	db, con, err := storage.DB()
	if err != nil {
		return
	}
	defer con.Close()

	db.Model(&model.Feed{}).Where("name = ?", feed.Name).Update("updated", feed.Updated)
}
