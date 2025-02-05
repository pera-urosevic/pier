package types

import "pier/api/monitor/database/model"

type MonitorData struct {
	Stats         []model.Stat         `json:"stats"`
	Notifications []model.Notification `json:"notifications"`
}
