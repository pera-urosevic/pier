package controller

import (
	"net/http"
	"pier/api/reader/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func DiscardArticle(r *gin.Engine) {
	r.PUT("/reader/discard/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := database.DiscardArticle(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
