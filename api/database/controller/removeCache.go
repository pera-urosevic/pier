package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func RemoveCache(r *gin.Engine) {
	r.DELETE("/database/cache/:key", func(c *gin.Context) {
		key := c.Param("key")

		err := database.RemoveCache(key)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
