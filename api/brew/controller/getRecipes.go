package controller

import (
	"net/http"
	"pier/api/brew/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetRecipes(r *gin.Engine) {
	r.GET("/brew/recipes", func(c *gin.Context) {
		recipes, err := database.GetRecipes()
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, recipes)
	})
}
