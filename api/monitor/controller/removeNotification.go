package controller

import (
	"net/http"
	"pier/api/monitor/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveNotification(r *gin.Engine) {
	r.DELETE("/monitor/notification/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.RemoveNotification(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
