package model

func (Stat) TableName() string {
	return "monitor"
}

type Stat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
