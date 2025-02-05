package database

import (
	"pier/api/monitor/database/model"
	"pier/storage"
)

func RemoveNotification(id int64) error {
	db, con, err := storage.DB()
	if err != nil {
		return err
	}
	defer con.Close()

	res := db.Delete(&model.Notification{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
