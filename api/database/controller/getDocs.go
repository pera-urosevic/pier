package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDocs(r *gin.Engine) {
	r.GET("/database/:database/collection/:collection", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		where := c.Query("where")
		if where == "" {
			where = "true"
		}

		order := c.Query("order")
		if order == "" {
			order = "id DESC"
		}

		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			offset = 0
		}

		docs, err := database.GetDocs(databaseName, collectionName, where, order, offset)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, docs)
	})
}
