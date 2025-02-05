package tvshows

import (
	"pier/api/tvshows/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetTVShows(r)
	controller.GetTVShow(r)
	controller.CreateTVShow(r)
	controller.UpdateTVShow(r)
	controller.RemoveTVShow(r)
	controller.TVMazeSearch(r)
	controller.TVMazeGet(r)
}
