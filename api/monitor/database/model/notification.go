package model

func (Notification) TableName() string {
	return "notify"
}

type Notification struct {
	ID        int64  `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Channel   string `json:"channel"`
	Topic     string `json:"topic"`
	Message   string `json:"message"`
}
