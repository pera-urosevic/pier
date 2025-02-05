package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func SetCache(r *gin.Engine) {
	r.POST("/database/cache/:key", func(c *gin.Context) {
		var value string

		key := c.Param("key")

		err := c.BindJSON(&value)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.SetCache(key, value)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
