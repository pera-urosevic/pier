package monitor

import (
	"pier/api/monitor/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetData(r)
	controller.RemoveNotification(r)
	controller.RemoveNotifications(r)
}
