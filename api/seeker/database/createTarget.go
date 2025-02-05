package database

import (
	"pier/api/seeker/database/model"
	"pier/storage"
)

func CreateTarget(target model.Target) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Create(&target)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
