package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func UpdateFacets(r *gin.Engine) {
	r.PUT("/database/:database/collection/:collection/facets", func(c *gin.Context) {
		var facets string

		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		err := c.BindJSON(&facets)
		if err != nil {
			lib.GinError(c, err, false)
		}

		err = database.UpdateFacets(databaseName, collectionName, facets)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, facets)
	})
}
