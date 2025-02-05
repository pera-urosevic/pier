package database

import (
	"pier/api/seeker/database/model"
	"pier/storage"
)

func RemoveTarget(title string) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Where("title = ?", title).Delete(&model.Target{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
