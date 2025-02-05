package controller

import (
	"net/http"
	"pier/api/colors/database"
	"pier/api/colors/database/model"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func SetColor(r *gin.Engine) {
	r.PUT("/colors", func(c *gin.Context) {
		var color = model.Color{}
		err := c.BindJSON(&color)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.SetColor(color)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
