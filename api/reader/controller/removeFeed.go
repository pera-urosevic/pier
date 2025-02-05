package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func RemoveFeed(r *gin.Engine) {
	r.DELETE("/reader/feed/:name", func(c *gin.Context) {
		name := c.Param("name")

		err := database.RemoveFeed(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
