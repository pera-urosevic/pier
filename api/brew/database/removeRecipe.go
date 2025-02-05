package database

import (
	"pier/api/brew/database/model"
	"pier/storage"
)

func RemoveRecipe(id int64) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Delete(&model.Recipe{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
