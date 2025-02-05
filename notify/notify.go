package notify

import (
	"fmt"
	"pier/api/monitor/database/model"
	"pier/storage"
	"time"
)

// internal

func notify(channel string, topic string, message string) {
	fmt.Println(channel, topic, message)

	db, con, err := storage.DB()
	if err != nil {
		return
	}
	defer con.Close()

	now := time.Now()

	res := db.Create(&model.Notification{
		Timestamp: now.Unix(),
		Channel:   channel,
		Topic:     topic,
		Message:   message,
	})
	if res.Error != nil {
		fmt.Println(res.Error)
	}

	db.Where("channel = 'info' AND timestamp < ?", now.Add(-1*time.Hour).Unix()).Delete(&model.Notification{})
	db.Where("channel = 'warn' AND timestamp < ?", now.Add(-1*time.Hour*24).Unix()).Delete(&model.Notification{})
	db.Where("channel = 'alert' AND timestamp < ?", now.Add(-1*time.Hour*24*7).Unix()).Delete(&model.Notification{})
}

// notifications by type

func Info(topic string, message string) {
	notify("info", topic, message)
}

func Warn(topic string, message string) {
	notify("warn", topic, message)
}

func Alert(topic string, message string) {
	notify("alert", topic, message)
}

// utils

func ErrorAlert(topic string, subtopic string, err error) {
	message := fmt.Sprintf("%s\n%v", subtopic, err)
	Alert(topic, message)
}

func ErrorWarn(topic string, subtopic string, err error) {
	message := fmt.Sprintf("%s\n%v", subtopic, err)
	Warn(topic, message)
}

func ErrorInfo(topic string, subtopic string, err error) {
	message := fmt.Sprintf("%s\n%v", subtopic, err)
	Info(topic, message)
}
