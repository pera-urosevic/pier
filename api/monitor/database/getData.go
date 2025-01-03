package database

import (
	"pier/api/monitor/types"
	"pier/storage"
)

func GetData() (types.MonitorData, error) {
	var monitorData = types.MonitorData{}

	db := storage.DB()

	rows, err := db.Query("SELECT * FROM `monitor`")
	if err != nil {
		return monitorData, err
	}

	var stats types.Stats
	for rows.Next() {
		var stat types.Stat
		err := rows.Scan(&stat.Key, &stat.Value)
		if err != nil {
			return monitorData, err
		}

		stats = append(stats, stat)
	}
	monitorData.Stats = stats

	rows, err = db.Query("SELECT `id`, `timestamp`, `channel`, `topic`, `message` FROM `notify` ORDER BY `timestamp` DESC")
	if err != nil {
		return monitorData, err
	}

	var notifications []types.Notification
	for rows.Next() {
		var notification types.Notification
		err := rows.Scan(&notification.ID, &notification.Timestamp, &notification.Channel, &notification.Topic, &notification.Message)
		if err != nil {
			return monitorData, err
		}

		notifications = append(notifications, notification)
	}
	monitorData.Notifications = notifications

	return monitorData, nil
}
