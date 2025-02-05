package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"pier/api/reader/database/model"
	"pier/notify"
	"pier/storage"

	"github.com/mmcdole/gofeed"
)

func Articles(feed *model.Feed, items []*gofeed.Item, threshold time.Duration) {
	db, con, err := storage.DB()
	if err != nil {
		return
	}
	defer con.Close()

	var articles []model.Article
	res := db.Where("feed_name = ?", feed.Name).Find(&articles)
	if res.Error != nil {
		notify.ErrorAlert("reader", "load articles", res.Error)
		return
	}

	articlesMap := map[string]model.Article{}
	for _, article := range articles {
		articlesMap[article.ID] = article
	}

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			notify.ErrorWarn("reader", "json marshal article data", err)
			continue
		}

		// get article datetime
		guid := item.GUID
		var dt time.Time
		datetime := item.UpdatedParsed
		if datetime == nil {
			datetime = item.PublishedParsed
		}
		if datetime == nil {
			dt = time.Unix(feed.Updated, 0)
		} else {
			dt = *datetime
		}
		// skip articles that are too old
		if dt.Before(time.Now().Add(-1 * threshold)) {
			continue
		}
		// skip articles that already exist
		id := fmt.Sprintf("%s|%s", dt, guid)
		_, exists := articlesMap[id]
		if exists {
			delete(articlesMap, id)
			continue
		}

		article := model.Article{
			ID:        id,
			FeedName:  feed.Name,
			Content:   string(data),
			Discarded: false,
		}

		res := db.Create(&article)
		if res.Error != nil {
			notify.ErrorAlert("reader", "create article", res.Error)
		}
	}

	// delete discarded articles no longer present in feed
	for articleId, article := range articles {
		if article.Discarded {
			res := db.Delete(&model.Article{}, articleId)
			if res.Error != nil {
				notify.ErrorAlert("reader", "delete article", res.Error)
			}
		}
	}
}
