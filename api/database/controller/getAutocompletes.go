package controller

import (
	"net/http"
	"pier/api/database/database"
	"pier/api/database/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func GetAutocompletes(r *gin.Engine) {
	r.POST("/database/:database/collection/:collection/autocompletes", func(c *gin.Context) {
		var autocompleteFields types.AutocompleteFields

		databaseName := c.Param("database")
		collectionName := c.Param("collection")

		err := c.BindJSON(&autocompleteFields)
		if err != nil {
			lib.GinError(c, err, false)
			return
		}

		autocompletes, err := database.GetAutocompletes(databaseName, collectionName, autocompleteFields)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, autocompletes)
	})
}
