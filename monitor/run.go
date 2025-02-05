package monitor

import (
	"fmt"
	"os"

	"pier/monitor/sensor"
)

func Run() {
	if os.Getenv("RUN_MONITOR") != "true" {
		return
	}

	fmt.Println("MONITOR")

	go sensor.Uptime()
	go sensor.Storage()
	go sensor.Temp()
	go sensor.Cpu()
	go sensor.Mem()
	go sensor.Swap()
	go sensor.Heartbeat()
}
