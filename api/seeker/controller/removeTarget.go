package controller

import (
	"net/http"
	"pier/api/seeker/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func RemoveTarget(r *gin.Engine) {
	r.DELETE("/seeker/target/:title", func(c *gin.Context) {
		title := c.Param("title")

		err := database.RemoveTarget(title)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		targets, err := database.GetTargets()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, targets)
	})
}
