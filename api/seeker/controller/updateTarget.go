package controller

import (
	"net/http"
	"pier/api/seeker/database"
	"pier/api/seeker/database/model"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func UpdateTarget(r *gin.Engine) {
	r.PUT("/seeker/target/:title", func(c *gin.Context) {
		title := c.Param("title")

		var target model.Target
		err := c.BindJSON(&target)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.UpdateTarget(title, target)
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
