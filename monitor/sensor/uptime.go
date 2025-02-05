package sensor

import (
	"time"

	"pier/monitor/db"
	"pier/notify"

	statsHost "github.com/shirou/gopsutil/v3/host"
)

const secondsMinute = 60
const secondsHour = secondsMinute * 60
const secondsDay = secondsHour * 24

func sensorUptime() {
	uptime, err := statsHost.Uptime()
	if err != nil {
		notify.ErrorAlert("monitor", "get uptime", err)
		return
	}
	days := uptime / secondsDay
	hours := (uptime % secondsDay) / secondsHour

	db.Set("uptime:days", days)
	db.Set("uptime:hours", hours)
}

func Uptime() {
	sensorUptime()
	tickerUptime := time.NewTicker(1 * time.Hour)
	for range tickerUptime.C {
		sensorUptime()
	}
}
