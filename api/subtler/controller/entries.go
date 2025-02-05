package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"pier/api/subtler/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func Entries(r *gin.Engine) {
	r.GET("/subtler/entries/*path", func(c *gin.Context) {
		path := os.Getenv("SUBTLER_ROOT")
		path = filepath.Join(path, c.Param("path"))
		dir, err := os.ReadDir(path)
		if err != nil {
			lib.GinError(c, err, true)
			return
		}

		entries := []types.Entry{}
		for _, f := range dir {
			entry := types.Entry{Name: f.Name(), Dir: f.IsDir()}
			entries = append(entries, entry)
		}

		c.JSON(http.StatusOK, entries)
	})
}
