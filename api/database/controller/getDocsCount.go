package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetDocsCount(r *gin.Engine) {
	r.GET("/database/:database/collection/:collection/count", func(c *gin.Context) {

		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		where := c.Query("where")
		where = strings.ReplaceAll(where, "|", "%")
		if where == "" {
			where = "true"
		}

		count, err := database.GetDocsCount(databaseName, collectionName, where)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, count)
	})
}
