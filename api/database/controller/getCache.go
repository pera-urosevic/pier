package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetCache(r *gin.Engine) {
	r.GET("/database/cache/:key", func(c *gin.Context) {
		key := c.Param("key")

		value, err := database.GetCache(key)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, value)
	})
}
