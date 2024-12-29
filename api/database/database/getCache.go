package database

import (
	"pier/storage"
)

func GetCache(key string) (string, error) {
	var value string

	db := storage.DB()

	err := db.QueryRow("SELECT `value` FROM `database_cache` WHERE `key`=?", key).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}