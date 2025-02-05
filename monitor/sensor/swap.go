package sensor

import (
	"time"

	"pier/monitor/db"
	"pier/notify"

	statsMem "github.com/shirou/gopsutil/v3/mem"
)

func swap() {
	swap, err := statsMem.SwapMemory()
	if err != nil {
		notify.ErrorAlert("monitor", "get swap", err)
		return
	}

	db.Set("swap:usage", swap.UsedPercent)
}

func Swap() {
	swap()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		swap()
	}
}
