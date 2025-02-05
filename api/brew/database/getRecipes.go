package database

import (
	"pier/api/brew/database/model"
	"pier/storage"
)

func GetRecipes() ([]model.Recipe, error) {
	var recipes []model.Recipe

	db, con, err := storage.DB()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	res := db.Order("name ASC").Find(&recipes)
	if res.Error != nil {
		return nil, res.Error
	}

	return recipes, nil
}
