package proxy

import (
	"pier/api/proxy/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.Proxy(r)
}
