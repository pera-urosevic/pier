package storage

import (
	"pier/notify"
	"pier/storage"
	"time"
)

func Cleanup(since time.Duration) {
	db := storage.DB()
	thresholdDate := time.Now().Add(-1 * since).Format("2006-01-02")
	_, err := db.Exec("DELETE FROM `reader_articles` WHERE `discarded` = 1 AND `id` < ?", thresholdDate)
	if err != nil {
		notify.Warn("reader", "Cleanup: "+err.Error())
	}
}
