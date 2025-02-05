package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func GetFeeds() ([]model.Feed, error) {
	var feeds = []model.Feed{}

	db, con, err := storage.DB()
	if err != nil {
		return feeds, err
	}
	defer con.Close()

	res := db.Find(&feeds)
	if res.Error != nil {
		return feeds, res.Error
	}

	return feeds, nil
}
