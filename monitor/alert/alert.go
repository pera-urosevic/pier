package alert

import (
	"fmt"
	"strconv"

	"pier/monitor/db"
	"pier/notify"
)

func Signal(subject string, threshold int, increase bool, text string) {
	if !increase {
		db.Set("alert:"+subject, 0)
		return
	}

	res := db.Get("alert:" + subject)
	if res == "" {
		res = "0"
	}
	score, err := strconv.Atoi(res)
	if err != nil {
		notify.ErrorAlert("monitor", "parse score", err)
		return
	}

	score = score + 1
	notify.Alert("monitor", fmt.Sprintf("%s - %d", subject, score))

	if score >= threshold {
		notify.Alert(subject, text)
		db.Set("alert:"+subject, 0)
		return
	}
	db.Set("alert"+subject, score)
}
