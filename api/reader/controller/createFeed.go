package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func CreateFeed(r *gin.Engine) {
	r.POST("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")

		feed, err := database.CreateFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, feed)
	})
}
