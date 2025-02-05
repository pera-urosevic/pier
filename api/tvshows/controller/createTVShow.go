package controller

import (
	"net/http"
	"pier/api/tvshows/database"
	"pier/api/tvshows/database/model"
	"pier/lib"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateTVShow(r *gin.Engine) {
	r.POST("/tvshows/tvshow", func(c *gin.Context) {
		var tvshow model.TVShow

		err := c.BindJSON(&tvshow)
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

		tvshow, err = database.CreateTVShow(tvshow)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, tvshow)
	})
}
