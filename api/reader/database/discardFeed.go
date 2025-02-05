package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func DiscardFeed(name string) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Model((&model.Article{})).Where("feed_name = ?", name).Update("Discarded", true)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
