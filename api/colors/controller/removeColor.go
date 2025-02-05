package controller

import (
	"net/http"
	"pier/api/colors/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func RemoveColor(r *gin.Engine) {
	r.DELETE("/colors/*name", func(c *gin.Context) {
		name := c.Param("name")[1:]

		err := database.RemoveColor(name)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
