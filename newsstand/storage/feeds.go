package storage

import (
	"fmt"
	"localhost/pier/database"
	"localhost/pier/newsstand/models"
	"localhost/pier/notify"
)

func Feeds() []*models.Feed {
	ids := feedIds()
	feeds := []*models.Feed{}

	db := database.Connect()
	for _, id := range ids {
		key := fmt.Sprintf("newsstand:%s:feed", id)
		hash, err := db.HGetAll(database.Ctx, key).Result()
		if err != nil {
			notify.ErrorAlert("newsstand", "get feeds", err)
			continue
		}
		feeds = append(feeds, &models.Feed{
			Id:       id,
			Url:      hash["url"],
			Disabled: hash["disabled"] == "true",
			Updated:  hash["updated"],
		})
	}

	return feeds
}
