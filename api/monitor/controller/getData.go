package controller

import (
	"net/http"
	"pier/api/monitor/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetData(r *gin.Engine) {
	r.GET("/monitor", func(c *gin.Context) {
		monitorData, err := database.GetData()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, monitorData)
	})
}
