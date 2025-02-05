package controller

import (
	"net/http"
	"pier/api/tvshows/database"
	"pier/api/tvshows/database/model"
	"pier/lib"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateTVShow(r *gin.Engine) {
	r.PUT("/tvshows/tvshow/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		var tvshow model.TVShow
		err = c.BindJSON(&tvshow)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		if tvshow.Image != nil {
			if strings.HasPrefix(*tvshow.Image, "http") {
				bytes, err := lib.Download(*tvshow.Image)
				if err != nil {
					lib.GinError(c, err, true)
					return
				}

				base64 := lib.Base64Encode(bytes)
				tvshow.Image = &base64
			}
		}

		err = database.UpdateTVShow(id, tvshow)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, tvshow)
	})
}
