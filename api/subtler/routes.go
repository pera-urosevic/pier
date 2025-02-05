package subtler

import (
	"pier/api/subtler/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.Entries(r)
	controller.Extract(r)
}
