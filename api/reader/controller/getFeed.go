package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetFeed(r *gin.Engine) {
	r.GET("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")

		feed, err := database.GetFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, feed)
	})
}
