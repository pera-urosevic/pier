package database

import (
	"pier/api/tvshows/database/model"
	"pier/storage"
)

func RemoveTVShow(id int64) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Delete(&model.TVShow{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
