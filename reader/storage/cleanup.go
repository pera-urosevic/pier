package storage

import (
	"pier/api/reader/database/model"
	"pier/notify"
	"pier/storage"
	"time"
)

func Cleanup(since time.Duration) {
	db, con, err := storage.DB()
	if err != nil {
		return
	}
	defer con.Close()

	thresholdDate := time.Now().Add(-1 * since).Format("2006-01-02")
	res := db.Where("discarded = ? AND id < ?", true, thresholdDate).Delete(&model.Article{})
	if res.Error != nil {
		notify.Warn("reader", "Cleanup: "+res.Error.Error())
	}
}
