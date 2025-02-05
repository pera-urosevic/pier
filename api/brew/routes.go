package brew

import (
	"pier/api/brew/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetRecipes(r)
	controller.CreateRecipe(r)
	controller.UpdateRecipe(r)
	controller.GetRecipe(r)
	controller.RemoveRecipe(r)
}
