package controller

import (
	"net/http"
	"pier/api/seeker/database"
	"pier/api/seeker/database/model"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func CreateTarget(r *gin.Engine) {
	r.POST("/seeker/target", func(c *gin.Context) {
		var target model.Target

		err := c.BindJSON(&target)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.CreateTarget(target)
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
