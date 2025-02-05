package database

import (
	"pier/api/colors/database/model"
	"pier/storage"
)

func RemoveColor(name string) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Where("name = ?", name).Delete(&model.Color{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
