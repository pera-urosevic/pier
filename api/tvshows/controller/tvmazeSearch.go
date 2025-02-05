package controller

import (
	"net/http"
	"pier/api/tvshows/tvmaze"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func TVMazeSearch(r *gin.Engine) {
	r.GET("/tvshows/tvmaze/search/:title", func(c *gin.Context) {
		title := c.Param("title")

		results, err := tvmaze.Search(title)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, results)
	})
}
