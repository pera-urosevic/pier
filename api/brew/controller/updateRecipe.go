package controller

import (
	"net/http"
	"pier/api/brew/database"
	"pier/api/brew/database/model"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func UpdateRecipe(r *gin.Engine) {
	r.PUT("/brew/recipe", func(c *gin.Context) {
		var recipe model.Recipe

		err := c.BindJSON(&recipe)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		recipe, err = database.UpdateRecipe(recipe)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, recipe)
	})
}
