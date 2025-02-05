package reader

import (
	"pier/api/reader/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetBundle(r)
	controller.GetFeed(r)
	controller.CreateFeed(r)
	controller.UpdateFeed(r)
	controller.RemoveFeed(r)
	controller.DiscardFeed(r)
	controller.DiscardArticle(r)
}
