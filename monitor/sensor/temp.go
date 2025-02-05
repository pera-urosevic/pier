package sensor

import (
	"fmt"
	"time"

	"pier/monitor/alert"
	"pier/monitor/db"
	"pier/notify"

	statsHost "github.com/shirou/gopsutil/v3/host"
)

func temp() {
	temps, err := statsHost.SensorsTemperatures()
	if err != nil {
		notify.ErrorAlert("monitor", "get temperature", err)
		return
	}

	db.Del("temp:%")
	for _, temp := range temps {
		alert.Signal("temp value", 1, temp.Temperature > 70.0, fmt.Sprintf("%s = %f", temp.SensorKey, temp.Temperature))
		db.Set("temp:"+temp.SensorKey, temp.Temperature)
	}
}

func Temp() {
	temp()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		temp()
	}
}
