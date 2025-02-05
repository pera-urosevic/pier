package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetDatabaseIndex(r *gin.Engine) {
	r.GET("/database/index", func(c *gin.Context) {
		databaseIndex, err := database.GetDatabaseIndex()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}
		c.JSON(http.StatusOK, databaseIndex)
	})
}
