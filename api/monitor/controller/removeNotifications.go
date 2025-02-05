package controller

import (
	"net/http"
	"pier/api/monitor/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func RemoveNotifications(r *gin.Engine) {
	r.DELETE("/monitor/notifications/:channel", func(c *gin.Context) {
		channel := c.Param("channel")

		err := database.RemoveNotifications(channel)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
