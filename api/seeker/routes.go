package seeker

import (
	"pier/api/seeker/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetTargets(r)
	controller.CreateTarget(r)
	controller.UpdateTarget(r)
	controller.RemoveTarget(r)
}
