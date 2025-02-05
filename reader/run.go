package reader

import (
	"fmt"
	"os"
	"time"

	"pier/notify"
	"pier/reader/net"
	"pier/reader/storage"
)

const INTERVAL = 20
const THRESHOLD = time.Hour * 24 * 28

func task() {
	feeds := storage.Feeds()
	for _, feed := range feeds {
		// skip disabled
		if feed.Disabled {
			continue
		}

		now := time.Now()
		// skip fresh
		then := time.Unix(feed.Updated, 0)
		diff := now.Sub(then).Minutes()
		if diff < INTERVAL {
			continue
		}

		// fetch feed
		res, err := net.Fetch(feed)
		if err != nil {
			notify.ErrorAlert("reader:fetch", "fetch feed "+feed.URL, err)
			continue
		}
		feed.Updated = now.Unix()
		status := fmt.Sprintf("fetched %s", feed.Name)
		notify.Info("reader", status)

		// store articles
		storage.Articles(feed, res.Items, THRESHOLD)

		// store meta
		storage.FeedUpdate(feed)
	}
	// cleanup old articles
	storage.Cleanup(THRESHOLD)
}

func check(lastRun time.Time) bool {
	reload := storage.Reload()
	if reload {
		notify.Info("reader", "reload")
		return true
	}
	now := time.Now()
	diff := now.Sub(lastRun).Minutes()
	return diff >= INTERVAL
}

func Run() {
	if os.Getenv("RUN_READER") != "true" {
		return
	}

	fmt.Println("READER")

	task()
	lastRun := time.Now()
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		run := check(lastRun)
		if run {
			task()
			lastRun = time.Now()
		}
	}
}
