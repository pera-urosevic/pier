package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func RemoveFeed(name string) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Where("name = ?", name).Delete(&model.Feed{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
