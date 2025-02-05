package controller

import (
	"net/http"
	"pier/api/brew/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecipe(r *gin.Engine) {
	r.GET("/brew/recipe/:id", func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		recipe, err := database.GetRecipe(id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, recipe)
	})
}
