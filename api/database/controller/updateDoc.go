package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/api/database/types"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateDoc(r *gin.Engine) {
	r.PUT("/database/:database/collection/:collection/doc/:id", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		var doc types.Doc
		err = c.BindJSON(&doc)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		err = database.UpdateDoc(databaseName, collectionName, id, doc)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		docNew, err := database.GetDoc(databaseName, collectionName, id)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, docNew)
	})
}
