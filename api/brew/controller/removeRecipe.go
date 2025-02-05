package controller

import (
	"net/http"
	"pier/api/brew/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveRecipe(r *gin.Engine) {
	r.DELETE("/brew/recipe/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.RemoveRecipe(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
