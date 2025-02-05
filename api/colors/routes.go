package colors

import (
	"pier/api/colors/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetColors(r)
	controller.SetColor(r)
	controller.RemoveColor(r)
}
