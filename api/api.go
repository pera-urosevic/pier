package api

import (
	"os"
	"pier/api/brew"
	"pier/api/colors"
	"pier/api/database"
	"pier/api/monitor"
	"pier/api/proxy"
	"pier/api/reader"
	"pier/api/seeker"
	"pier/api/subtler"
	"pier/api/tvshows"
	"pier/lib"
	"pier/notify"
)

func Run() {
	if os.Getenv("RUN_API") != "true" {
		return
	}

	host := os.Getenv("API_HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "55555"
	}
	addr := host + ":" + port

	notify.Info("api", addr)

	router := lib.CustomGin()

	brew.Routes(router)
	colors.Routes(router)
	database.Routes(router)
	monitor.Routes(router)
	proxy.Routes(router)
	reader.Routes(router)
	seeker.Routes(router)
	subtler.Routes(router)
	tvshows.Routes(router)

	router.Run(addr)
}
