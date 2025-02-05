package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDoc(r *gin.Engine) {
	r.GET("/database/:database/collection/:collection/doc/:id", func(c *gin.Context) {

		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		doc, err := database.GetDoc(databaseName, collectionName, id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, doc)
	})
}
