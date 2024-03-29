package storage

import (
	"localhost/pier/database"
	"regexp"
)

func feedIds() []string {
	ids := []string{}

	db := database.Connect()
	keys, err := db.Keys(database.Ctx, "newsstand:*:feed").Result()
	if err != nil {
		return []string{}
	}

	re := regexp.MustCompile(`newsstand:(.*):feed`)
	for _, key := range keys {
		matches := re.FindStringSubmatch(key)
		ids = append(ids, matches[1])
	}

	return ids
}
