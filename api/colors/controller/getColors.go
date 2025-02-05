package controller

import (
	"net/http"
	"pier/api/colors/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetColors(r *gin.Engine) {
	r.GET("/colors", func(c *gin.Context) {
		colors, err := database.GetColors()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, colors)
	})
}
