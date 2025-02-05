package database

import (
	"pier/api/reader/database/model"
	"pier/storage"
)

func GetArticles() ([]model.Article, error) {
	var articles = []model.Article{}

	db, con, err := storage.DB()
	if err != nil {
		return articles, err
	}
	defer con.Close()

	res := db.Where("discarded = ?", false).Find(&articles)
	if res.Error != nil {
		return articles, res.Error
	}

	return articles, nil
}
