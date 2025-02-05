package controller

import (
	"net/http"
	"pier/api/tvshows/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetTVShows(r *gin.Engine) {
	r.GET("/tvshows", func(c *gin.Context) {
		tvshows, err := database.GetTVShows()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, tvshows)
	})
}
