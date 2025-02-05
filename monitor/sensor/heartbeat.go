package sensor

import (
	"time"

	"pier/monitor/db"
)

func heartbeat() {
	timestamp := time.Now().UnixNano() / 1000

	db.Set("heartbeat:timestamp", timestamp)
}

func Heartbeat() {
	heartbeat()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		heartbeat()
	}
}
