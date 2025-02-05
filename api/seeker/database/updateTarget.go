package database

import (
	"pier/api/seeker/database/model"
	"pier/storage"
)

func UpdateTarget(title string, target model.Target) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Where("title = ?", title).Save(&target)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
