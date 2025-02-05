package database

import (
	"pier/api/tvshows/database/model"
	"pier/storage"
)

func GetTVShows() ([]model.TVShow, error) {
	var tvshows []model.TVShow

	db, con, err := storage.DB()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	res := db.Find(&tvshows)
	if res.Error != nil {
		return nil, res.Error
	}

	return tvshows, nil
}
