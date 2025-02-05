package controller

import (
	"net/http"
	"pier/api/tvshows/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveTVShow(r *gin.Engine) {
	r.DELETE("/tvshows/tvshow/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.RemoveTVShow(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
