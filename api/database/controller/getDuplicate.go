package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDuplicate(r *gin.Engine) {
	r.GET("/database/:database/collection/:collection/duplicate/:id/:name", func(c *gin.Context) {
		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		name := c.Param("name")
		exists, err := database.GetDuplicate(databaseName, collectionName, id, name)
		if err != nil {
			lib.GinError(c, err, true)
		}

		c.JSON(http.StatusOK, exists)
	})
}
