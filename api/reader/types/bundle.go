package types

import "pier/api/reader/database/model"

type Bundle struct {
	Feeds    []model.Feed    `json:"feeds"`
	Articles []model.Article `json:"articles"`
}
