package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetFacets(r *gin.Engine) {
	r.GET("/database/:database/collection/:collection/facets", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		facets, err := database.GetFacets(databaseName, collectionName)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, facets)
	})
}
