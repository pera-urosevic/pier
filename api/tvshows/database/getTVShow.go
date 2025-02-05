package database

import (
	"pier/api/tvshows/database/model"
	"pier/storage"
)

func GetTVShow(id int64) (model.TVShow, error) {
	var tvshow model.TVShow

	db, con, err := storage.DB()
	if err != nil {
		return tvshow, err
	}
	defer con.Close()

	res := db.First(&tvshow, id)
	if res.Error != nil {
		return tvshow, res.Error
	}

	return tvshow, nil
}
