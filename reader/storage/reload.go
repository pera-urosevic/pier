package storage

import "pier/storage"

func Reload() bool {
	db, con, err := storage.DB()
	if err != nil {
		return false
	}
	defer con.Close()

	var count int64
	res := db.Where("disabled = 0 AND updated = 0").Count(&count)
	if res.Error != nil {
		return false
	}

	return count > 0
}
