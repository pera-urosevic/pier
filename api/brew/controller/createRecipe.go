package controller

import (
	"net/http"
	"pier/api/brew/database"
	"pier/api/brew/database/model"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func CreateRecipe(r *gin.Engine) {
	r.POST("/brew/recipe", func(c *gin.Context) {
		var recipe model.Recipe

		err := c.BindJSON(&recipe)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		recipe, err = database.CreateRecipe(recipe)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, recipe)
	})
}
