package sensor

import (
	"fmt"
	"strings"
	"time"

	"pier/monitor/alert"
	"pier/monitor/db"
	"pier/notify"

	statsDisk "github.com/shirou/gopsutil/v3/disk"
)

func storage() {
	partitions, err := statsDisk.Partitions(true)
	if err != nil {
		notify.ErrorAlert("monitor", "get partitions", err)
		return
	}

	db.Del("storage:%")
	for _, partition := range partitions {
		if strings.HasPrefix(partition.Mountpoint, "/run") {
			continue
		}
		usage, err := statsDisk.Usage(partition.Mountpoint)
		if err != nil {
			notify.ErrorAlert("monitor", "get partition usage ("+partition.Mountpoint+")", err)
			continue
		}
		if usage.UsedPercent < 1 {
			continue
		}
		alert.Signal("storage usage", 1, usage.UsedPercent > 90.0, fmt.Sprintf("%s - %f", partition.Mountpoint, usage.UsedPercent))
		db.Set("storage:"+partition.Mountpoint, usage.UsedPercent)
	}
}

func Storage() {
	storage()
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		storage()
	}
}
