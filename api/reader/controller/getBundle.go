package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/api/reader/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetBundle(r *gin.Engine) {
	r.GET("/reader/bundles", func(c *gin.Context) {
		feeds, err := database.GetFeeds()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		articles, err := database.GetArticles()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		bundle := types.Bundle{Feeds: feeds, Articles: articles}
		c.JSON(http.StatusOK, bundle)
	})
}
