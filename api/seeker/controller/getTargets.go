package controller

import (
	"net/http"
	"pier/api/seeker/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetTargets(r *gin.Engine) {
	r.GET("/seeker/targets", func(c *gin.Context) {
		targets, err := database.GetTargets()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, targets)
	})
}
