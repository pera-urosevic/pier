package database

import (
	"pier/api/tvshows/database/model"
	"pier/storage"
)

func UpdateTVShow(id int64, record model.TVShow) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Save(&record)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
