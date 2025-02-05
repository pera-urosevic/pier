package model

import "time"

type Tabler interface {
	TableName() string
}

func (TVShow) TableName() string {
	return "tvshows"
}

type TVShow struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Status    string     `json:"status"`
	Premiered *time.Time `json:"premiered"`
	TVMaze    int64      `json:"tvmaze" gorm:"column:tvmaze"`
	IMDB      *string    `json:"imdb"`
	Website   *string    `json:"website"`
	Updated   time.Time  `json:"updated"`
	Episodes  string     `json:"episodes"`
	Watching  int64      `json:"watching"`
	Image     *string    `json:"image"`
	Stream    *string    `json:"stream"`
	Runtime   *int64     `json:"runtime"`
}
