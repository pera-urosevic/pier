package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func CreateFeed(name string) (model.Feed, error) {
	var feed = model.Feed{
		Name:     name,
		URL:      "",
		Web:      "",
		Icon:     "",
		Tokens:   "{}",
		Disabled: true,
		Updated:  0,
		Style:    "",
	}

	db, con, err := storage.DB()
	if err != nil {
		return feed, err
	}
	defer con.Close()

	res := db.Create(&feed)
	if res.Error != nil {
		return feed, res.Error
	}

	return feed, nil
}
