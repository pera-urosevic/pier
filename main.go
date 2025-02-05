package main

import (
	"os"

	"pier/api"
	"pier/env"
	"pier/monitor"
	"pier/notify"
	"pier/reader"
	// "net/http"
	// _ "net/http/pprof"
)

func main() {
	env.Load()

	notify.Info("pier", "starting")

	go monitor.Run()
	go reader.Run()
	go api.Run()

	// go func() {
	// 	http.ListenAndServe("0.0.0.0:9999", nil)
	// }()

	done := make(chan os.Signal, 1)
	<-done

	notify.Info("pier", "ending")
}
