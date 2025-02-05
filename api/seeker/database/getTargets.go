package database

import (
	"pier/api/seeker/database/model"
	"pier/storage"
)

func GetTargets() ([]model.Target, error) {
	var targets = []model.Target{}

	db, con, err := storage.DB()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	res := db.Order("`release` asc").Order("`title` asc").Find(&targets)
	if res.Error != nil {
		return nil, res.Error
	}

	return targets, nil
}
