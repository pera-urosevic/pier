package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/api/database/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func CreateDoc(r *gin.Engine) {
	r.POST("/database/:database/collection/:collection/doc", func(c *gin.Context) {
		var doc types.Doc

		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		err := c.BindJSON(&doc)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		id, err := database.CreateDoc(databaseName, collectionName, doc)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, id)
	})
}
