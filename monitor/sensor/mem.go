package sensor

import (
	"fmt"
	"time"

	"pier/monitor/alert"
	"pier/monitor/db"
	"pier/notify"

	statsMem "github.com/shirou/gopsutil/v3/mem"
)

func mem() {
	vm, err := statsMem.VirtualMemory()
	if err != nil {
		notify.ErrorAlert("monitor", "get virtual memory", err)
		return
	}

	alert.Signal("mem usage", 2, vm.UsedPercent > 90.0, fmt.Sprintf("%f", vm.UsedPercent))
	db.Set("mem:usage", vm.UsedPercent)
}

func Mem() {
	mem()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		mem()
	}
}
