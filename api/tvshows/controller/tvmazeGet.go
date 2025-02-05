package controller

import (
	"net/http"
	"pier/api/tvshows/tvmaze"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TVMazeGet(r *gin.Engine) {
	r.POST("/tvshows/tvmaze/get/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
		}

		result, err := tvmaze.Get(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
