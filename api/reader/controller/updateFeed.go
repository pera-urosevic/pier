package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/api/reader/database/model"
	"pier/lib"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateFeed(r *gin.Engine) {
	r.PUT("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")

		var feed model.Feed
		err := c.BindJSON(&feed)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		if strings.HasPrefix(feed.Icon, "http") {
			bytes, err := lib.Download(feed.Icon)
			if err != nil {
				lib.GinError(c, err, true)
				return
			}

			feed.Icon = lib.Base64Encode(bytes)
		}
		feed, err = database.UpdateFeed(name, feed)

		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, feed)
	})
}
