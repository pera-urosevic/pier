package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func DiscardFeed(r *gin.Engine) {
	r.PUT("/reader/discard/feed/:name", func(c *gin.Context) {
		name := c.Param("name")

		err := database.DiscardFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
